package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "A simple command-line tool",
	Long:  "MyCLI is a simple and powerful command-line tool built with Cobra.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello! Welcome to My CLI 🚀")
	},
}

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet someone with a friendly message",
	Long:  "The greet command takes a name and prints a welcome message.",
	Run: func(cmd *cobra.Command, args []string) {
		name := "Guest"
		if len(args) > 0 {
			name = args[0]
		}
		fmt.Println("Welcome 🚀", name)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CLI Version: 1.0.0")
	},
}

func Execute() {
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(versionCmd)
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}