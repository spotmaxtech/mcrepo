package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "repository management",
	Long:  "repository management",
}

var RepoList = &cobra.Command{
	Use:   "list",
	Short: "repository list",
	Long:  "repository list",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infof("repository list command run")
	},
}

func init() {
	RepoCmd.AddCommand(RepoList)
}