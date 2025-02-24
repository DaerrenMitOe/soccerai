package dataset

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const repoUrl string = "https://github.com/statsbomb/open-data"
const dest string = "dataset"

func downloadGitHubRepo(repoUrl string, dest string) {
	cmd := exec.Command("git", "clone", repoUrl, dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Repository cloned successfully!")
	}
}

func checkGitHubRepoDownloaded(dest string) bool {
	if _, err := os.Stat(filepath.Join(dest, "README.md")); !os.IsNotExist(err) {
		fmt.Println("Repository already exists.")
		return true
	}
	return false
}

func deleteFolder(dest string) {
	err := os.RemoveAll(dest)
	if err != nil {
		fmt.Println("Error deleting folder:", err)
	} else {
		fmt.Println("Folder deleted successfully!")
	}
}

func Download() {
	if !checkGitHubRepoDownloaded(dest) {
		downloadGitHubRepo(repoUrl, dest)
	}
	deleteFolder("dataset/.git")
}
