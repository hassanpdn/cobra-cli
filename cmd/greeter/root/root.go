package root

import (
	"github.com/spf13/cobra"
	"fmt"
	"cobra-cli/cmd/greeter/greet"
)

func RootCommand() *cobra.Command{
	cmd := &cobra.Command{
		Use:   "root",
		Short: "A simple command-line tool",
		Long:  "MyCLI is a simple and powerful command-line tool built with Cobra.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello! Welcome to My CLI 🚀")
		},
	}

	cmd.AddCommand(greet.GreetCommand())
	return cmd
}