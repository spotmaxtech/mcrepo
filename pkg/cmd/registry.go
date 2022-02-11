package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RegistryCmd = &cobra.Command{
	Use:   "registry",
	Short: "registry management",
	Long:  "registry management",
}

var RegistryList = &cobra.Command{
	Use:   "list",
	Short: "registry list",
	Long:  "registry list",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infof("registry list command run")
	},
}

func init() {
	RegistryCmd.AddCommand(RegistryList)
}
