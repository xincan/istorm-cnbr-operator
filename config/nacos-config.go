/***
 * nacos配置中心配置
 * @author JiangXincan
 * @date 2022/2/16 16:45
 **/
package config

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NacosRegisterInstance() {

	// Nacos服务配置
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: viper.GetString("nacos.client.ip"),
			Port:   viper.GetUint64("nacos.client.port"),
		},
	}
	// Nacos客户端配置
	clientConfig := constant.ClientConfig{
		TimeoutMs:           viper.GetUint64("nacos.timeoutMs"),
		NotLoadCacheAtStart: true,
		NamespaceId:         viper.GetString("nacos.client.namespaceId"),
		Username:            viper.GetString("nacos.client.username"),
		Password:            viper.GetString("nacos.client.password"),
		CacheDir:            viper.GetString("nacos.cacheDir"),
		LogDir:              viper.GetString("nacos.logDir"),
		LogRollingConfig: &constant.ClientLogRollingConfig{
			MaxSize: 1,
		},
	}

	nacosConfigInstance := setNacosNamingClient(serverConfig, clientConfig)
	boo, err := nacosConfigInstance.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          viper.GetString("nacos.client.ip"),
		Port:        viper.GetUint64("nacos.client.port"),
		ServiceName: viper.GetString("service.name"),
		Weight:      10,
		GroupName:   viper.GetString("nacos.client.group"),
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})
	if !boo {
		logrus.WithField("fail", err).Info("Nacos服务注册失败")
		return
	}
	logrus.WithField("success", boo).Info("Nacos服务注册成功")
}

/***
 * 设置注册中心
 * @author JiangXincan
 * @date 2022/2/16 16:41
 * @return null
 **/
func setNacosNamingClient(serverConfig []constant.ServerConfig, clientConfig constant.ClientConfig) naming_client.INamingClient {
	// 配置租户注册中心
	var namingClient, err = clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfig,
	})
	if err != nil {
		logrus.WithField("fail", err).Error("配置Nacos租户注册中心失败")
		return nil
	}

	setNacosConfigClient(serverConfig, clientConfig)
	return namingClient
}

/***
 * 设置配置中心
 * @author JiangXincan
 * @date 2022/2/16 16:41
 **/
func setNacosConfigClient(serverConfig []constant.ServerConfig, clientConfig constant.ClientConfig) {
	configClient, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfig,
	})
	if err != nil {
		logrus.WithField("fail", err).Error("配置Nacos配置中心失败")
	}
	setNacosConfigLister(configClient)
}

/***
 * 监听、更新、获取Nacos配置中心数据
 * @param configClient nacos配置中心客户端对象
 * @author JiangXincan
 * @date 2022/2/16 16:39
 **/
func setNacosConfigLister(configClient config_client.IConfigClient) {
	configClient.ListenConfig(vo.ConfigParam{
		DataId: viper.GetString("nacos.config.dataId"),
		Group:  viper.GetString("nacos.client.group"),
		OnChange: func(namespace, group, dataId, data string) {
			logrus.WithFields(logrus.Fields{
				"dataId": dataId,
				"group":  group,
				"data":   data,
			}).Info("更新Nacos配置中心数据")
		},
	})

	data, err := configClient.GetConfig(vo.ConfigParam{
		DataId: viper.GetString("nacos.config.dataId"),
		Group:  viper.GetString("nacos.client.group"),
	})

	if err != nil {
		logrus.WithField("fail", err).Error("获取Nacos配置中心数据失败")
	}

	logrus.WithFields(logrus.Fields{
		"DataId": viper.GetString("nacos.config.dataId"),
		"Group":  viper.GetString("nacos.client.group"),
		"data":   data,
	}).Info("监听Nacos配置中心数据")

}
