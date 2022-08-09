package cmd

import (
	"github.com/spf13/cobra"
)

var (
	limit int
)

// listCmd represents the list command
func newListCmd(collection, task *cobra.Command) *cobra.Command {
	c := &cobra.Command{
		Use:   "list",
		Short: "Perform the list action",
		Long:  `Get the list of tasks to do`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	c.AddCommand(collection)
	c.AddCommand(task)

	return c
}
