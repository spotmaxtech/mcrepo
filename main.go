package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	rootCmd := &cobra.Command{
		Use:   "mcrepo",
		Short: "maxcloud multi repo management tool",
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
