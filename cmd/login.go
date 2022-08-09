/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
func newLoginCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Login to your google account",
		Long:  `Login to your google account through a auth code.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("login called")
		},
	}
}
