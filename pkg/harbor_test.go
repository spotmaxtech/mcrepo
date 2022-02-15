package pkg

import (
	"context"
	"fmt"
	"github.com/spotmaxtech/gokit"
	"testing"
)
import "github.com/mittwald/goharbor-client/v5/apiv1"

func TestHarbor(t *testing.T) {
	host, err := apiv1.NewRESTClientForHost("", "", "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(gokit.PrettifyYaml(host))
	registry, err := host.GetRegistry(context.Background(), "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(gokit.PrettifyYaml(registry))
	projects, err := host.ListProjects(context.Background(), "*")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(gokit.PrettifyYaml(projects))
}