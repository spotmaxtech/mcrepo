package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spotmaxtech/mcrepo/pkg/cmd"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	rootCmd := &cobra.Command{
		Use:   "mcrepo",
		Short: "maxcloud multi repo management tool",
		Long:  "maxcloud multi repo management tool",
	}

	rootCmd.AddCommand(cmd.RegistryCmd)
	rootCmd.AddCommand(cmd.RepoCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
