package pkg

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestAliyunRegistry(t *testing.T) {
	InitConfig()
	InitAliyunRegistryMap(GMcrepoConfig.Aliyun)
	for name, registry := range AliyunRegistryMap {
		logrus.Infof("registry name: %s", name)
		registry.ListInstance()
	}
}
