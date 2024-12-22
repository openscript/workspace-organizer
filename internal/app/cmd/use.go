package cmd

import (
	"fmt"

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
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
