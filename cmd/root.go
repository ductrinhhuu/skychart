package cmd

import (
	"github.com/spf13/cobra"
)

var (
	Root = &cobra.Command{
		Use: "skychart [command]",
	}
)

func init() {
	Root.AddCommand(ServerCmd)
}

func Execute() error {
	return Root.Execute()
}
