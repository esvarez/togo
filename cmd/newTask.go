/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"github.com/esvarez/togo/config"
	"github.com/esvarez/togo/internal/entity"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var (
	name        string
	description string
)

func newTaskCmd(ctx context.Context, t taskUseCase, app *config.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "Create new task",
		Long:  `Create new task.`,
		Run: func(cmd *cobra.Command, args []string) {
			if taskList == "" {
				taskList = app.DefaultList
			}
			now := time.Now()
			task := &entity.Task{
				Name:        name,
				Description: description,
				DueDate:     now.Format("2006-01-02T00:00:00.000Z"),
			}
			id, err := t.Create(ctx, taskList, task)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("Task created with id: %s\n", id)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the task")
	cmd.Flags().StringVarP(&description, "description", "d", "", "Description of the task")
	cmd.Flags().StringVarP(&taskList, "taskList", "l", "", "Name of the task list")

	cmd.MarkFlagRequired("name")
	return cmd
}
