package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hadenlabs/pagerduty-exporter/config"
	pkg "github.com/hadenlabs/pagerduty-exporter/pkg/version"
)

var (
	version bool
)

func init() {
	if err := config.InitializeViper(); err != nil {
		panic(err)
	}
	cobra.OnInitialize()
	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "show current version of CLI")
}

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "LuisMayta PagerdutyExporter",
	Short: "LuisMayta go server",
	Long:  `PagerdutyExporter Tool`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return config.InitializeViper()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			return printVersion()
		}
		return cmd.Help()
	},
}

func printVersion() error {
	fmt.Println("Version: ", pkg.Version)
	return nil
}

// Execute command cmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
