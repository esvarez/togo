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
	cmd := &cobra.Command{
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
	cmd.Flags().StringVarP(&taskList, "taskList", "l", "", "Name of the task list")
	return cmd
}
