package main

import (
	"fmt"
	"cobra-cli/cmd/greeter/root"
)

func main() {
	rootCmd := root.RootCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
