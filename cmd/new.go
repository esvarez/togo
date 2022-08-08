package cmd

import (
	"github.com/spf13/cobra"
)

var (
	name        string
	description string
)

// newCmd represents the new command
func createNewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "Create a new task",
		Long:  `Create a new task.`,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
}

func initAddNewCmd(n *cobra.Command) {
	n.Flags().StringVarP(&name, "name", "n", "", "Name of the task")
	n.Flags().StringVarP(&description, "description", "d", "", "Description of the task")

	n.MarkFlagRequired("name")
}
