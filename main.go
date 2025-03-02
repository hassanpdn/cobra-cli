package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main(){
	rootCmd := &cobra.Command{
		Use: "root",
		Short: "A simple CLI tool",
		Long:  "This is a simple command-line tool built with Cobra.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello! Welcome to My CLI 🚀")
		},
	}

	greetCmd := &cobra.Command{
		Use: "greet",
		Short: "A simple CLI tool",
		Long:  "This is a simple command-line tool built with Cobra.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome 🚀", args[0])
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the CLI version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("CLI Version: 1.0.0")
		},
	}

	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}