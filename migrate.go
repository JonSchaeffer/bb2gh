package main

import (
	"log"
	"fmt"
	"os/exec"
	"encoding/json"
	"os"
)

func migrate(bitbucketOrg string, bitbucketRepos []string, githubOrg string) {
	tempDir := "./tmp/repos"
	mkdirCmd := exec.Command("mkdir", "-p", tempDir)
	err := mkdirCmd.Run()
	if err != nil {
		log.Fatalf("Failed to create temp directory: %v", err)	
	}

	existingRepos := getExistingGithubRepos(githubOrg)


	for _, repo := range bitbucketRepos {
		fmt.Printf("Migrating bitbucket:%s/%s to github:%s/%s\n", bitbucketOrg, repo, githubOrg, repo)

		tempRepo := fmt.Sprintf("%s/%s.git", tempDir, repo)

		//Clone the repo from Bitbucket
		cmd := exec.Command("git", "-C", tempDir, "clone", "--mirror", "--origin", "bitbucket", fmt.Sprintf("git@bitbucket.org:%s/%s.git", bitbucketOrg, repo))
		err := cmd.Run()
		if err != nil {
			os.RemoveAll(tempDir)
			log.Fatalf("Failed to clone Bitbucket/%s: %v", repo, err)
		}

		// Create a new repo in Github (if it doesn't already exist)
		if !contains(existingRepos, repo) {
			cmd = exec.Command("gh", "repo", "create", fmt.Sprintf("%s/%s", githubOrg, repo), "--private")
			err = cmd.Run()
			if err != nil {
				os.RemoveAll(tempDir)
				log.Fatalf("Failed to create Github/%s: %v", repo, err)
			}
		}

		// Push all refs to Github
		cmd = exec.Command("git", "-C", tempRepo, "push", "--mirror", fmt.Sprintf("https://github.com/%s/%s.git", githubOrg, repo))
		err = cmd.Run()
		if err != nil {
			os.RemoveAll(tempDir)
			log.Fatalf("Failed to push to Github: %v", err)
		}

		// Remove the temporary local repository.
		err = os.RemoveAll(tempRepo)
		if err != nil {
			log.Fatalf("Failed to remove temp repo: %v", err)
		}

		fmt.Printf("Completed migration for %s\n", repo)
			
	}
	os.RemoveAll("./tmp")
	fmt.Println("Migration complete!")
	
}

// contains checks if a slice contains a string.
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func getExistingGithubRepos (githubOrg string) []string {

	// Get a list of existing repos in the GitHub org.
	cmd := exec.Command("gh", "repo", "list", githubOrg, "--json", "name")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to list GitHub repos: %v", err)
	}

	var repos []struct {
		Name string `json:"name"`
	}
	err = json.Unmarshal(output, &repos)
	if err != nil {
		log.Fatalf("Failed to parse GitHub repos: %v", err)
	}

	existingRepos := make([]string, len(repos))
	for i, repo := range repos {
		existingRepos[i] = repo.Name
	}

	return existingRepos
}
