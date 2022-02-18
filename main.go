// Title  main.go
// Description  程序主入口
// Author  jiangxincan@hatech.com.cn  2021/1/22 11:36
// update  jiangxincan@hatech.com.cn  2021/1/22 11:36
package main

import (
	"github.com/sirupsen/logrus"
	"hatech.com.cn/istorm-cnbr-operator/config"
	_ "hatech.com.cn/istorm-cnbr-operator/docs"
	"hatech.com.cn/istorm-cnbr-operator/logs"
	"net/http"
)

// @Title 计算器API
// @version 1.0
// @Description 提供加减乘除工具
// @termsOfService http://www.hatech.com.cn
// @contact.name hatech
// @contact.url http://www.hatech.com.cn
// @contact.email hatech@hatech.com.cn
// @host localhost:8080
// @BasePath /
// @Auth      			jiangxincan@hatech.com.cn       时间（2021/1/22 11:36）
// @Return    			router        	*mux.Router     "返回路由对象"
func main() {

	// 配置本地yaml读取
	config.SetViperYaml()

	// 配置远程yaml读取
	setViperRemoteYaml := config.SetViperRemoteYaml()

	// 配置日志
	logrus.WithField("file", setViperRemoteYaml.GetString("service.log")).Info("日志信息存放地址")
	logrus.AddHook(logs.NewHook(setViperRemoteYaml.GetString("service.log")))

	// 配置注册中心
	config.NacosRegisterInstance()

	logrus.WithField("port", setViperRemoteYaml.GetUint64("service.port")).Info("服务启动端口")
	_ = http.ListenAndServe(":"+setViperRemoteYaml.GetString("service.port"), Router())
}
