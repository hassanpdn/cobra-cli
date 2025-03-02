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
			fmt.Println("Welcome Hassan! 🚀")
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

	rootCmd.AddCommand(greetCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}