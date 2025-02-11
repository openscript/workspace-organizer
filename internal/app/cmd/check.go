package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:     "check",
	Short:   "Check the environment",
	Long:    "Check the environment for the workspace-organizer.",
	Args:    cobra.ExactArgs(0),
	Example: "check",
	Run: func(cmd *cobra.Command, args []string) {
		checkForGit()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func checkForGit() {
	_, err := exec.LookPath("git")
	if err != nil {
		fmt.Println("❌ git is not installed or not found in PATH")
	} else {
		fmt.Println("✅ git is available.")
	}
}
