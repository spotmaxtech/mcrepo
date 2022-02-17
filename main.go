package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spotmaxtech/mcrepo/pkg"
	"github.com/spotmaxtech/mcrepo/pkg/cmd"
)

func main() {
	logrus.SetLevel(logrus.ErrorLevel)
	rootCmd := &cobra.Command{
		Use:   "mcrepo",
		Short: "maxcloud multi repo management tool",
		Long:  "maxcloud multi repo management tool",
	}

	// config and init
	pkg.InitConfig()

	pkg.InitAliyunRegistryMap(pkg.GMcrepoConfig.Aliyun)
	if pkg.GMcrepoConfig.CurrentContext.Platform == "aliyun" {
		pkg.CurrentRegistry = pkg.AliyunRegistryMap[pkg.GMcrepoConfig.CurrentContext.RegistryName]
	}

	pkg.InitHarborRegistryMap(pkg.GMcrepoConfig.Harbor)
	if pkg.GMcrepoConfig.CurrentContext.Platform == "harbor" {
		pkg.CurrentRegistry = pkg.HarborRegistryMap[pkg.GMcrepoConfig.CurrentContext.RegistryName]
	}

	rootCmd.AddCommand(cmd.RegistryCmd)
	rootCmd.AddCommand(cmd.RepoCmd)
	rootCmd.AddCommand(cmd.ImageCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
