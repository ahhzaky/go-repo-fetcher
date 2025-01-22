package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type Repository struct {
	Name string `json:"name"`
}

func main() {
	// Input username GitHub
	var username string
	fmt.Print("Insert username GitHub: ")
	fmt.Scanln(&username)

	// Create a folder with the username
	if err := os.Mkdir(username, os.ModePerm); err != nil {
		if !os.IsExist(err) {
			fmt.Printf("Error creating folder: %v\n", err)
			return
		}
	}

	// Change working directory to the created folder
	if err := os.Chdir(username); err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		return
	}

	// Fetch repository data
	repos, err := fetchRepos(username)
	if err != nil {
		fmt.Printf("Error fetching repositories: %v\n", err)
		return
	}

	if len(repos) == 0 {
		fmt.Println("No public repositories found.")
		return
	}

	// Download each repository
	for _, repo := range repos {
		fmt.Printf("Cloning repository: %s\n", repo.Name)
		if err := cloneRepo(username, repo.Name); err != nil {
			fmt.Printf("[Failed to clone repository %s: %v]\n\n", repo.Name, err)
		} else {
			fmt.Printf("[Successfully cloned repository: %s]\n\n", repo.Name)
		}
	}
}

func fetchRepos(username string) ([]Repository, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch repositories, status code: %d", resp.StatusCode)
	}

	var repos []Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}

func cloneRepo(username, repoName string) error {
	repoURL := fmt.Sprintf("https://github.com/%s/%s.git", username, repoName)
	cmd := exec.Command("git", "clone", repoURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
