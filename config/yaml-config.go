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
	filName      = "istorm-cnbr-operator-dev"
	suffix       = "yaml"
	provider     = "nacos"
)

// Title    			SetViperYaml
// Description   		设置本地读取yaml
// Auth      			jiangxincan@hatech.com.cn       			时间（2022/02/17 11:36）
// Return    			naming_client.INamingClient     			"无返回值"
func SetViperYaml() {
	viper.SetConfigName(filName)
	viper.AddConfigPath(".")
	viper.SetConfigType(suffix)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error("Viper插件加载本地配置数据失败")
	}
	logrus.Info("Viper插件加载本地配置数据成功")
}

// Title    			SetViperRemoteYaml
// Description   		设置远程读取yaml
// Auth      			jiangxincan@hatech.com.cn       			时间（2022/02/17 11:36）
// Return    			*viper.Viper     							"Viper yaml 解析器实例"
func SetViperRemoteYaml() *viper.Viper {
	configViper := viper.New()
	runtimeViper := configViper
	runtimeViper.SetConfigFile("./" + filName + "." + suffix)
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
	err := remoteViper.AddRemoteProvider(provider, configViper.GetString("nacos.client.ip"), "")
	remoteViper.SetConfigType(suffix)
	err = remoteViper.ReadRemoteConfig()
	if err == nil {
		configViper = remoteViper
		logrus.WithFields(logrus.Fields{"config": configViper.Get("config"), "service": configViper.Get("service"), "nacos": configViper.Get("nacos")}).Info("Viper 远程获取 Nacos 配置中心配置数据")
		remoteProvider := remote.NewRemoteProvider(suffix)
		watchRemoteConfigOnChannel := remoteProvider.WatchRemoteConfigOnChannel(configViper)
		go func(rc <-chan bool) {
			for {
				<-rc
				logrus.WithFields(logrus.Fields{
					"config":  configViper.Get("config"),
					"service": configViper.Get("service"),
					"nacos":   configViper.Get("nacos"),
				}).Info("Viper 更新远程 Nacos 配置中心配置数据")
			}
		}(watchRemoteConfigOnChannel)
	}
	RemoteConfig = configViper
	return configViper
}
