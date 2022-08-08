package cmd

import (
	"github.com/spf13/cobra"
)

var (
	limit int
)

// listCmd represents the list command
func newListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Perform the list action",
		Long:  `Get the list of tasks to do`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
}

func initListCmd(list *cobra.Command, collection *cobra.Command, task *cobra.Command) {
	list.AddCommand(collection)
	list.AddCommand(task)
}
