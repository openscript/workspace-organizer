package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"workspace-organizer/internal/app/config"

	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:     "use <url>",
	Short:   "Use a project",
	Long:    "Use a project by cloning, if it doesn't exists, and changing to its directory.",
	Args:    cobra.ExactArgs(1),
	Example: "use https://github.com/openscript/workspace-organizer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("URL:", args[0])

		url := args[0]
		workspacePath := config.GetWorkspaceFolder()

		hoster, org, repo := parseURL(url)
		projectPath := filepath.Join(workspacePath, hoster, org, repo)

		if _, err := os.Stat(projectPath); os.IsNotExist(err) {
			cloneProject(url, projectPath)
		}

		changeDirectory(projectPath)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}

func parseURL(url string) (hoster, org, repo string) {
	if strings.HasPrefix(url, "http") {
		parts := strings.Split(url, "/")
		hoster = parts[2]
		org = parts[3]
		repo = strings.TrimSuffix(parts[4], ".git")
	} else if strings.HasPrefix(url, "git@") {
		parts := strings.Split(url, ":")
		hoster = strings.Split(parts[0], "@")[1]
		pathParts := strings.Split(parts[1], "/")
		org = pathParts[0]
		repo = strings.TrimSuffix(pathParts[1], ".git")
	}
	return
}

func cloneProject(url, projectPath string) {
	cmd := execCommand("git", "clone", url, projectPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to clone the project: %v", err)
	}
}

func changeDirectory(path string) {
	if err := os.Chdir(path); err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}
	fmt.Printf("Changed directory to %s\n", path)
}
