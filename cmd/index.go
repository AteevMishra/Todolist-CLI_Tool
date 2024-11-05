package cmd

import "github.com/spf13/cobra"

// RootCmd is the root command for the CLI
var RootCmd = cobra.Command{
	Use: "todo-CLI",
	Short: "A CLI tool for managing todos",
}

func Execute() error{
	return RootCmd.Execute()
}