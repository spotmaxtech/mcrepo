package pkg

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spotmaxtech/goharbor-client/v5/apiv2"
	"github.com/spotmaxtech/goharbor-client/v5/apiv2/pkg/config"
	"github.com/spotmaxtech/gokit"
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
		fmt.Println(err.Error())
	}
	fmt.Println(gokit.PrettifyYaml(projects))
	repositories, err := r.Client.ListRepositories(context.Background(), "official-website")
	if err != nil {
		return
	}
	fmt.Println(gokit.PrettifyYaml(repositories))
}

func (r *HarborRegistry) ListRepoTag(repoName string) {
}
