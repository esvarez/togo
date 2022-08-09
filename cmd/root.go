package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/esvarez/togo/config"
	"github.com/esvarez/togo/internal/entity"
	"github.com/esvarez/togo/internal/usecase"
	"github.com/esvarez/togo/internal/usecase/repo"
	"github.com/esvarez/togo/pkg/google"

	"github.com/spf13/cobra"
)

type taskUseCase interface {
	Create(context.Context, string, *entity.Task) (string, error)
	List(context.Context, string, int) ([]*entity.Task, error)
}

type listUseCase interface {
	List(context.Context) ([]*entity.TaskList, error)
	Create(context.Context, *entity.TaskList) (string, error)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "togo",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	srv := google.NewService()

	taskRepo := repo.NewTaskRepo(srv)
	listRepo := repo.NewListRepo(srv)

	taskUC := usecase.NewTask(taskRepo)
	listUC := usecase.NewList(listRepo)

	collectionListCmd := newCollectionListCmd(ctx, listUC)
	taskListCmd := newTaskListCmd(ctx, taskUC, &cfg.App)

	nwListCmd := newCreateListCmd(ctx, listUC)
	nwTaskCmd := newTaskCmd(ctx, taskUC, &cfg.App)

	listCmd := newListCmd(collectionListCmd, taskListCmd)
	newCmd := createNewCmd(nwListCmd, nwTaskCmd)

	configCmd := newConfigCmd(cfg)

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(newCmd)

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.togo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
