package cmd

import (
	"github.com/esvarez/togo/config"
	"github.com/spf13/cobra"
)

var defaultList string

func newConfigCmd(conf *config.Configuration) *cobra.Command {
	c := &cobra.Command{
		Use:   "config",
		Short: "Configure the application",
		Long:  `Configure the application.`,
		Run: func(cmd *cobra.Command, args []string) {
			conf.App.DefaultList = defaultList
			config.SaveConfiguration(conf)
		},
	}

	c.Flags().StringVarP(&defaultList, "default_list", "d", "", "ID of list")
	c.MarkFlagRequired("default_list")
	return c
}
