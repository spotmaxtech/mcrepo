package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spotmaxtech/mcrepo/pkg"
)

var ImageCmd = &cobra.Command{
	Use:   "image",
	Short: "image management",
	Long:  "image management",
}

var imageList = &cobra.Command{
	Use:   "list",
	Short: "image list",
	Long:  "image list",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.CurrentRegistry.ListRepoTag(args[0])
		logrus.Debugf("image list command run")
	},
}

func init() {
	ImageCmd.AddCommand(imageList)
}
