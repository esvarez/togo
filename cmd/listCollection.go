/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

// taskCollectionListCmd represents the taskCollectionList command
func newCollectionListCmd(ctx context.Context, l listUseCase) *cobra.Command {
	return &cobra.Command{
		Use:   "collection",
		Short: "Get the collection of tasks to do",
		Long:  `Get the collection of tasks to do`,
		Run: func(cmd *cobra.Command, args []string) {
			lists, err := l.List(ctx)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
			for _, list := range lists {
				fmt.Println(list)
			}
		},
	}
}
