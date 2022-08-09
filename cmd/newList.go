/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"github.com/esvarez/togo/internal/entity"
	"os"

	"github.com/spf13/cobra"
)

var listName string

func newCreateListCmd(ctx context.Context, l listUseCase) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Create a new list",
		Long:  `Create a new list.`,
		Run: func(cmd *cobra.Command, args []string) {
			list := &entity.TaskList{
				Title: listName,
			}
			id, err := l.Create(ctx, list)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("List created with id: %s\n", id)
		},
	}

	cmd.Flags().StringVarP(&listName, "listName", "n", "", "Name of the list")
	cmd.MarkFlagRequired("listName")

	return cmd
}
