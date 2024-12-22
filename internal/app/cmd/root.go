package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wo",
	Short: "Workspace Organizer",
	Long: `Workspace Organizer

Workspace Organizer is a tool to keep your workspace nice and tidy, while
helping you to navigate and manage your projects. No more messy repositories
scattered all over the place.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Run() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops: %v\n", err)
		os.Exit(1)
	}
}
