package pkg

import (
	"testing"
)

func TestHarborRegistry_ListRepo(t *testing.T) {
	InitConfig()
	InitHarborRegistryMap(GMcrepoConfig.Harbor)
	HarborRegistryMap["second"].ListRepo()
}

func TestHarborRegistry_ListRepoTag(t *testing.T) {
	InitConfig()
	InitHarborRegistryMap(GMcrepoConfig.Harbor)
	HarborRegistryMap["second"].ListRepoTag("official-website/cms_server")
}