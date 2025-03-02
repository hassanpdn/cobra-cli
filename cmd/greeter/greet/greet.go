package greet

import (
	"github.com/spf13/cobra"
	"fmt"
)

func GreetCommand() *cobra.Command {
	return &cobra.Command{
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
}