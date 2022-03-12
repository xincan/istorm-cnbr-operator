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

// Title    			NacosRegisterInstance
// Description   		配置nacos注册实例
// Auth      			jiangxincan@hatech.com.cn       			时间（2022/02/17 11:36）
// Return    			nil     									"无返回值"
func NacosRegisterInstance() {
	// Nacos服务配置
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: viper.GetString("nacos.ip"),
			Port:   viper.GetUint64("nacos.port"),
		},
	}
	// Nacos客户端配置
	clientConfig := constant.ClientConfig{
		TimeoutMs:           viper.GetUint64("nacos.timeoutMs"),
		NotLoadCacheAtStart: true,
		NamespaceId:         viper.GetString("nacos.client.namespaceId"),
		Username:            viper.GetString("nacos.username"),
		Password:            viper.GetString("nacos.password"),
		CacheDir:            viper.GetString("nacos.cacheDir"),
		LogDir:              viper.GetString("nacos.logDir"),
		LogRollingConfig: &constant.ClientLogRollingConfig{
			MaxSize: 1,
		},
	}
	nacosConfigInstance := setNacosNamingClient(serverConfig, clientConfig)
	boo, err := nacosConfigInstance.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          viper.GetString("nacos.ip"),
		Port:        viper.GetUint64("nacos.port"),
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

// Title    			setNacosNamingClient
// Description   		设置Nacos命名空间客户端
// Auth      			jiangxincan@hatech.com.cn       			时间（2022/02/17 11:36）
// Return    			naming_client.INamingClient     			"返回Nacos命名空间客户端接口"
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
	// checked:[nacos, viper]  有两个选项，nacos：利用nacos原生提供config配置同步，viper：利用第三方工具进行config同步。默认使用第三方工具
	if viper.GetString("nacos.config.checked") == "nacos" {
		setNacosConfigClient(serverConfig, clientConfig)
	}
	return namingClient
}

// Title    			setNacosConfigClient
// Description   		设置Nacos配置中心客户端
// Auth      			jiangxincan@hatech.com.cn       			时间（2022/02/17 11:36）
// Return    			nil     									"无返回值"
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

// Title    			setNacosConfigLister
// Description   		设置Nacos配置中心数据监听
// Auth      			jiangxincan@hatech.com.cn       			时间（2022/02/17 11:36）
// Return    			naming_client.INamingClient     			"无返回值"
func setNacosConfigLister(configClient config_client.IConfigClient) {
	configClient.ListenConfig(vo.ConfigParam{
		DataId: viper.GetString("nacos.config.dataId"),
		Group:  viper.GetString("nacos.config.group"),
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
		Group:  viper.GetString("nacos.config.group"),
	})

	if err != nil {
		logrus.WithField("fail", err).Error("获取Nacos配置中心数据失败")
	}

	logrus.WithFields(logrus.Fields{
		"DataId": viper.GetString("nacos.config.dataId"),
		"Group":  viper.GetString("nacos.config.group"),
		"data":   data,
	}).Info("监听Nacos配置中心数据")

}
