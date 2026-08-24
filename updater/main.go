package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	login        = "willpinha"
	minRepoStars = 10
)

func main() {
	root := flag.String("root", "..", "path to the repository root")
	flag.Parse()

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN environment variable is not set")
	}

	client := newClient(token)

	repos, err := client.famousRepos()
	if err != nil {
		log.Fatalf("fetch repositories: %v", err)
	}
	pullRequests, err := client.createdPullRequests()
	if err != nil {
		log.Fatalf("fetch pull requests: %v", err)
	}
	issues, err := client.participatedIssues()
	if err != nil {
		log.Fatalf("fetch issues: %v", err)
	}
	discussions, err := client.participatedDiscussions()
	if err != nil {
		log.Fatalf("fetch discussions: %v", err)
	}

	files := map[string]string{
		"README.md":             renderReadme(time.Now(), repos, pullRequests, issues, discussions),
		"docs/pull-requests.md": renderSeeAll("All pull requests I created", pullRequests),
		"docs/issues.md":        renderSeeAll("All issues I participated in", issues),
		"docs/discussions.md":   renderSeeAll("All discussions I participated in", discussions),
	}
	for name, content := range files {
		path := filepath.Join(*root, name)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			log.Fatalf("create directory for %s: %v", name, err)
		}
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			log.Fatalf("write %s: %v", name, err)
		}
		log.Printf("wrote %s", name)
	}
}
