package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var wrapperCmd = &cobra.Command{
	Use:     "wrapper <install|uninstall>",
	Short:   "Manages the wrapper scripts",
	Long:    "Install or uninstall the wrapper scripts to use the workspace-organizer commands.",
	Args:    cobra.ExactArgs(1),
	Example: "wrapper install",
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "install":
			installWrappers()
		case "uninstall":
			uninstallWrappers()
		default:
			fmt.Println("Invalid argument. Use 'install' or 'uninstall'.")
		}
	},
}

func init() {
	rootCmd.AddCommand(wrapperCmd)
}

func installWrappers() {
	programPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting the program path.")
		os.Exit(1)
	}
	fmt.Printf("Program started with argument: %s\n", programPath)
	fmt.Println("Wrapper installed.")
}

func uninstallWrappers() {
	fmt.Println("Wrapper uninstalled.")
}
