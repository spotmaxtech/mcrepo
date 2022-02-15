package pkg

import (
	"fmt"
	"github.com/spf13/viper"
)

var GMcrepoConfig = McrepoConfig{}

type McrepoConfig struct {
	Aliyun         []*AliyunConfig `mapstructure:"aliyun"`
	CurrentContext CurrentContext  `mapstructure:"current_context"`
}

type CurrentContext struct {
	Platform     string `mapstructure:"platform"`
	RegistryName string `mapstructure:"registry_name"`
}

type AliyunConfig struct {
	Name         string `mapstructure:"name"`
	AccessKeyId  string `mapstructure:"access_key_id"`
	AccessSecret string `mapstructure:"access_secret"`
	Endpoint     string `mapstructure:"endpoint"`
	InstanceName string `mapstructure:"instance_name"`
}

func InitConfig() {
	viper.SetConfigName(".mcrepo") // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/")  // call multiple times to add many search paths
	viper.AddConfigPath(".")       // optionally look for config in the working directory
	err := viper.ReadInConfig()    // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	err = viper.Unmarshal(&GMcrepoConfig)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
