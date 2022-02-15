package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spotmaxtech/gokit"
	"github.com/spotmaxtech/mcrepo/pkg"
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

var RegistryAll = &cobra.Command{
	Use:   "all",
	Short: "registry all",
	Long:  "registry all",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(gokit.PrettifyYaml(pkg.GMcrepoConfig))
	},
}

func init() {
	RegistryCmd.AddCommand(RegistryAll)
	RegistryCmd.AddCommand(RegistryList)
}
