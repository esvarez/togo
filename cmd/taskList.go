package cmd

import (
	"context"
	"fmt"

	"github.com/esvarez/togo/config"
	"github.com/spf13/cobra"
)

var (
	taskList string
)

// taskListCmd represents the taskList command
func newTaskListCmd(ctx context.Context, t taskUseCase, app *config.App) *cobra.Command {
	return &cobra.Command{
		Use:   "task",
		Short: "List all tasks",
		Long:  `List all tasks.`,
		Run: func(cmd *cobra.Command, args []string) {
			if taskList == "" {
				taskList = app.DefaultList
			}

			tasks, err := t.List(ctx, taskList, limit)
			if err != nil {
				fmt.Errorf(err.Error())
			}

			for _, task := range tasks {
				fmt.Println(task)
			}
		},
	}
}

func initTaskList(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&taskList, "taskList", "l", "", "Name of the task list")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
