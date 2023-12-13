package util

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"strings"
)

var Config = new(GConfig)

type GConfig struct {
	Handle string `yaml:"handle"`
}

func ReadConfigAndSetup(configName string, GConfig interface{}) error {
	v, err := ReadConfig(configName)
	if err != nil {
		return err
	}
	err = v.Unmarshal(GConfig)
	if err != nil {
		return err
	}
	return nil
}

func ReadConfigToBytes(configName string) ([]byte, error) {
	v, err := ReadConfig(configName)
	if err != nil {
		return nil, err
	}
	settings := v.AllSettings()
	bytes, err := yaml.Marshal(settings)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func ReadConfig(configName string) (*viper.Viper, error) {
	// yaml
	vconfig := viper.New()
	//表示 先预加载匹配的环境变量
	vconfig.AutomaticEnv()
	//设置环境变量分割符，将点号和横杠替换为下划线
	vconfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	// 设置配置文件的路径
	vconfig.SetConfigName(configName)
	vconfig.AddConfigPath(".")
	vconfig.SetConfigType("yaml")

	// 读取配置文件
	err := vconfig.ReadInConfig()
	if err != nil {
		_ = fmt.Errorf("service error: %s\n", err)
		return nil, err
	}
	return vconfig, nil
}
