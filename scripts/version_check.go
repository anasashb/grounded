// Small CI script to make sure manifests from .version-files.json and git
// tag are in sync
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const configPath = ".version-files.json"

var pinned = regexp.MustCompile(`^\d+\.\d+\.\d+$`)

// expects all manifest files to be semantically versioned
func versionOf(path string) string {
	raw, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var manifest struct {
		Version string `json:"version"`
	}
	if err := json.Unmarshal(raw, &manifest); err != nil {
		log.Fatalf("%s: %v", path, err)
	}
	if !pinned.MatchString(manifest.Version) {
		log.Fatalf("%s: want X.Y.Z, got %q", path, manifest.Version)
	}
	return manifest.Version
}

func main() {
	log.SetFlags(0)

	raw, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	var config struct {
		Files []string `json:"files"`
	}
	if err := json.Unmarshal(raw, &config); err != nil {
		log.Fatalf("%s: %v", configPath, err)
	}
	files := config.Files

	first := versionOf(files[0])
	for _, path := range files[1:] {
		if version := versionOf(path); version != first {
			log.Fatalf("%s is %s, but %s is %s", path, version, files[0], first)
		}
	}

	if os.Getenv("GITHUB_REF_TYPE") == "tag" {
		tag := os.Getenv("GITHUB_REF_NAME")
		if strings.TrimPrefix(tag, "v") != first {
			log.Fatalf("tag %s does not match version %s", tag, first)
		}
	}

	fmt.Printf("%d manifests pinned at %s\n", len(files), first)
}
