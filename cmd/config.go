package cmd

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"

	"github.com/esvarez/togo/config"
	"github.com/spf13/cobra"
)

var defaultList string

const path = "./config/config.yaml"

func newConfigCmd(conf *config.Configuration) *cobra.Command {
	c := &cobra.Command{
		Use:   "config",
		Short: "Configure the application",
		Long:  `Configure the application.`,
		Run: func(cmd *cobra.Command, args []string) {
			conf.App.DefaultList = defaultList
			f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
			if err != nil {
				log.Fatalf("Unable to cache oauth token: %v", err)
			}
			defer f.Close()
			bytes, err := yaml.Marshal(conf)
			if err != nil {
				log.Fatalf("Unable to cache oauth token: %v", err)
			}

			_, err = f.Write(bytes)
			if err != nil {
				log.Fatalf("Unable to cache oauth token: %v", err)
			}
		},
	}

	c.Flags().StringVarP(&defaultList, "default_list", "d", "", "ID of list")
	c.MarkFlagRequired("default_list")
	return c
}
