package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
func createNewCmd(list, task *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new task",
		Long:  `Create a new task.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("new called")
		},
	}
	cmd.AddCommand(list)
	cmd.AddCommand(task)

	return cmd
}
