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

var (
	RemoteConfig *viper.Viper
	yamlFile     = "istorm-cnbr-operator-dev"
)

func SetViperYaml() {
	viper.SetConfigName(yamlFile)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error("Viper插件加载本地配置数据失败")
	}
	logrus.Info("Viper插件加载本地配置数据成功")
}

func SetViperRemoteYaml() *viper.Viper {
	configViper := viper.New()
	runtimeViper := configViper
	runtimeViper.SetConfigFile("./" + yamlFile + ".yaml")
	_ = runtimeViper.ReadInConfig()
	remote.SetOptions(&remote.Option{
		Url:         configViper.GetString("nacos.client.ip"),
		Port:        configViper.GetUint64("nacos.client.port"),
		NamespaceId: configViper.GetString("nacos.client.namespaceId"),
		GroupName:   configViper.GetString("nacos.client.group"),
		Config: remote.Config{
			DataId: configViper.GetString("nacos.config.dataId"),
		},
		Auth: nil,
	})
	remoteViper := viper.New()
	err := remoteViper.AddRemoteProvider("nacos", configViper.GetString("nacos.client.ip"), "")
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
				}).Info("Viper获取远程Nacos配置中心配置数据")

			}
		}(respChan)
	}
	RemoteConfig = configViper
	return configViper
}
