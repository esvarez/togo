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

func newDeleteListCmd(ctx context.Context, l listUseCase, app *config.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Delete a list",
		Long:  `Delete a list.`,
		Run: func(cmd *cobra.Command, args []string) {
			if listName == "" {
				listName = app.DefaultList
			}
			err := l.Delete(ctx, listName)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("List deleted: %s\n", listName)
		},
	}

	cmd.Flags().StringVarP(&listName, "listName", "l", "", "Name of the list")
	cmd.MarkFlagRequired("listName")
	return cmd
}
