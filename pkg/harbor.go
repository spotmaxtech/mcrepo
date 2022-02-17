package pkg

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spotmaxtech/goharbor-client/v5/apiv2"
	"github.com/spotmaxtech/goharbor-client/v5/apiv2/pkg/config"
	"strings"
)

var HarborRegistryMap map[string]*HarborRegistry

type HarborRegistry struct {
	Client *apiv2.RESTClient
}

func InitHarborRegistryMap(configList []*HarborConfig) {
	HarborRegistryMap = make(map[string]*HarborRegistry)
	for _, cfg := range configList {
		HarborRegistryMap[cfg.Name] = newHarborRegistry(cfg)
		logrus.Debugf("load harbor registry [%s]", cfg.Name)
	}
	return
}

func newHarborRegistry(harborConfig *HarborConfig) *HarborRegistry {
	clientHost, err := apiv2.NewRESTClientForHost(harborConfig.Api,
		harborConfig.Username, harborConfig.Password, &config.Options{})
	if err != nil {
		fmt.Println(err.Error())
	}
	return &HarborRegistry{
		Client: clientHost,
	}
}

func (r *HarborRegistry) ListRepo() {
	projects, err := r.Client.ListProjects(context.Background(), "")
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	for _, p := range projects {
		repositories, err := r.Client.ListRepositories(context.Background(), p.Name)
		if err != nil {
			logrus.Fatalln(err.Error())
		}
		for _, repo := range repositories {
			fmt.Println(repo.Name)
		}
	}
}

func (r *HarborRegistry) ListRepoTag(repoName string) {
	items := strings.Split(repoName, "/")
	if len(items) != 2 {
		logrus.Fatalln("please input project/repo format")
	}
	project, repo := items[0], items[1]

	artifacts, err := r.Client.ListArtifacts(context.Background(), project, repo)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, a := range artifacts {
		fmt.Printf("%s ", a.Digest)
		for _, t := range a.Tags {
			fmt.Printf("%s ", t.Name)
		}
		fmt.Println()
	}

}
