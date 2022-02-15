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
	Client       *cr20181201.Client
	InstanceName string
	InstanceId   string
}

func InitAliyunRegistryMap(configList []*AliyunConfig) {
	AliyunRegistryMap = make(map[string]*AliyunRegistry)
	for _, config := range configList {
		AliyunRegistryMap[config.Name] = newAliyunRegistry(config)
		//logrus.Infof("load aliyun registry [%s]", config.Name)
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

	request := &cr20181201.ListInstanceRequest{
		InstanceName: tea.String(config.InstanceName),
	}
	response, err := client.ListInstance(request)
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	var instanceId string
	for _, i := range response.Body.Instances {
		if *i.InstanceName == config.InstanceName {
			instanceId = *i.InstanceId
		}
	}
	return &AliyunRegistry{
		Client:       client,
		InstanceName: config.InstanceName,
		InstanceId:   instanceId,
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
func (r *AliyunRegistry) ListRepo() {
	request := &cr20181201.ListRepositoryRequest{
		InstanceId: tea.String(r.InstanceId),
		RepoStatus: tea.String("NORMAL"),
	}
	response, err := r.Client.ListRepository(request)
	if err != nil {
		logrus.Fatalln(err.Error())
	}
	for _, repo := range response.Body.Repositories {
		fmt.Printf("%s/%s\n", *repo.RepoNamespaceName, *repo.RepoName)
	}
}

func (r *AliyunRegistry) ListRepoTag(repoId string) {
	request := &cr20181201.ListRepoTagRequest{
		InstanceId: tea.String(r.InstanceId),
		RepoId:     tea.String(repoId),
	}
	response, err := r.Client.ListRepoTag(request)
	if err != nil {
		logrus.Fatalln(err.Error())
	}
	fmt.Println(gokit.PrettifyYaml(response))
	for _, image := range response.Body.Images {
		fmt.Println(*image.ImageId, *image.Tag)
	}
}
