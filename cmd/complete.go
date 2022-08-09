/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"github.com/esvarez/togo/config"
	"github.com/spf13/cobra"
	"os"
)

var taskID string

func newCompleteCmd(cxt context.Context, t taskUseCase, app *config.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "complete",
		Short: "Complete a task",
		Long:  `Complete a task.`,
		Run: func(cmd *cobra.Command, args []string) {
			if taskList == "" {
				taskList = app.DefaultList
			}
			err := t.Complete(cxt, taskList, taskID)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("Task completed: %s\n", taskID)
		},
	}

	cmd.Flags().StringVarP(&taskID, "taskID", "i", "", "ID of the task")
	cmd.Flags().StringVarP(&taskList, "taskList", "l", "", "Name of the task list")

	cmd.MarkFlagRequired("taskID")
	return cmd
}
