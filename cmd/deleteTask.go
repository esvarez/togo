/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"github.com/esvarez/togo/config"
	"os"

	"github.com/spf13/cobra"
)

var (
	listID string
)

func newDeleteTaskCmd(ctx context.Context, t taskUseCase, app *config.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "Delete task",
		Long:  `Delete task.`,
		Run: func(cmd *cobra.Command, args []string) {
			if listID == "" {
				listID = app.DefaultList
			}
			err := t.Delete(ctx, listID, taskID)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("Task deleted: %s\n", taskID)
		},
	}

	cmd.Flags().StringVarP(&listID, "listID", "l", "", "ID of the list")
	cmd.Flags().StringVarP(&taskID, "taskID", "i", "", "ID of the task")

	cmd.MarkFlagRequired("taskID")

	return cmd
}
