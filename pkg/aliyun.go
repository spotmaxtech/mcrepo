package pkg

import (
	"fmt"
	cr20181201 "github.com/alibabacloud-go/cr-20181201/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/sirupsen/logrus"
	"github.com/spotmaxtech/gokit"
)

var AliyunRegistryMap map[string]*AliyunRegistry

type AliyunRegistry struct {
	Client *cr20181201.Client
}

func InitAliyunRegistryMap(configList []*AliyunConfig) {
	AliyunRegistryMap = make(map[string]*AliyunRegistry)
	for _, config := range configList {
		AliyunRegistryMap[config.Name] = newAliyunRegistry(config)
		logrus.Infof("load aliyun registry [%s]", config.Name)
	}
	return
}

func newAliyunRegistry(config *AliyunConfig) *AliyunRegistry {
	apiConfig := &openapi.Config{
		AccessKeyId:     tea.String(config.AccessKeyId),
		AccessKeySecret: tea.String(config.AccessSecret),
		Endpoint:        tea.String(config.Endpoint),
	}
	client, err := cr20181201.NewClient(apiConfig)
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	return &AliyunRegistry{
		Client: client,
	}
}

func (r *AliyunRegistry) ListInstance() {
	request := &cr20181201.ListInstanceRequest{}
	response, err := r.Client.ListInstance(request)
	if err != nil {
		logrus.Fatalln(err.Error())
	}
	fmt.Println(gokit.PrettifyYaml(response))
}
func (r *AliyunRegistry) ListRepo(instanceId string) {
	request := &cr20181201.ListRepositoryRequest{
		InstanceId: tea.String(instanceId),
		RepoStatus: tea.String("NORMAL"),
	}
	response, err := r.Client.ListRepository(request)
	if err != nil {
		logrus.Fatalln(err.Error())
	}
	for _, repo := range response.Body.Repositories {
		fmt.Println(*repo.RepoName, *repo.RepoNamespaceName)
	}
}

func (r *AliyunRegistry) ListImage() {

}
