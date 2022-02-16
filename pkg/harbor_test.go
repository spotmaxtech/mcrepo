package pkg

import (
	"testing"
)

func TestHarborRegistry_ListRepo(t *testing.T) {
	InitConfig()
	InitHarborRegistryMap(GMcrepoConfig.Harbor)
	HarborRegistryMap["second"].ListRepo()
}