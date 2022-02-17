// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  xincan  2022/2/16 16:53
// @Update  xincan  2022/2/16 16:53
package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	remote "github.com/yoyofxteam/nacos-viper-remote"
)

var RemoteConfig *viper.Viper

func SetYaml() {
	viper.SetConfigName("test-gokit-dev")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error("Viper插件加载本地配置数据失败")
	}
	logrus.Info("Viper插件加载本地配置数据成功")
}

func SetRemoteConfig() {
	configViper := viper.New()
	runtimeViper := configViper
	runtimeViper.SetConfigFile("./test-gokit-dev.yaml")
	_ = runtimeViper.ReadInConfig()
	remote.SetOptions(&remote.Option{
		Url:         "192.168.1.81",
		Port:        31150,
		NamespaceId: "xincan-dev-0001",
		GroupName:   "dev_group",
		Config: remote.Config{
			DataId: "test-gokit-dev.yaml",
		},
		Auth: nil,
	})
	remoteViper := viper.New()
	err := remoteViper.AddRemoteProvider("nacos", "192.168.1.81", "")
	remoteViper.SetConfigType("yaml")
	err = remoteViper.ReadRemoteConfig()
	if err == nil {
		configViper = remoteViper
		logrus.WithField("success", true).Error("Viper 使用远程Nacos配置中心数据")
		provider := remote.NewRemoteProvider("yaml")
		respChan := provider.WatchRemoteConfigOnChannel(configViper)
		go func(rc <-chan bool) {
			for {
				<-rc
				logrus.WithFields(logrus.Fields{
					"config":  configViper.Get("config"),
					"service": configViper.Get("service"),
					"nacos":   configViper.Get("nacos"),
				}).Error("Viper remote get data config")

			}
		}(respChan)
	}
	RemoteConfig = configViper

}
