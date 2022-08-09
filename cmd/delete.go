package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
func newDeleteCmd(list, task *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete element",
		Long:  `Delete element.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("delete called")
		},
	}

	cmd.AddCommand(list)
	cmd.AddCommand(task)

	return cmd
}
