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

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCollectionListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskCollectionListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
