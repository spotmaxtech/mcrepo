package pkg

import (
	cr20181201  "github.com/alibabacloud-go/cr-20181201/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"log"
	"testing"
)

func CreateClient (accessKeyId *string, accessKeySecret *string) (_result *cr20181201.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId: accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("")
	_result = &cr20181201.Client{}
	_result, _err = cr20181201.NewClient(config)
	return _result, _err
}

func TestAliyunCR(t *testing.T) {
	client, _err := CreateClient(tea.String(""), tea.String(""))
	if _err != nil {
		log.Fatalf(_err.Error())
	}

	getInstanceRequest := &cr20181201.GetInstanceRequest{
		InstanceId: tea.String(""),
	}
	// 复制代码运行请自行打印 API 的返回值
	_, _err = client.GetInstance(getInstanceRequest)
	if _err != nil {
		log.Fatalf(_err.Error())
	}
}

