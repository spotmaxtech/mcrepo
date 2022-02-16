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
	for _, config := range configList {
		HarborRegistryMap[config.Name] = newHarborRegistry(config)
		logrus.Debugf("load harbor registry [%s]", config.Name)
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
	registries, err := r.Client.ListRegistries(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(gokit.PrettifyYaml(registries))
}

func (r *HarborRegistry) ListRepoTag(repoName string) {
}
