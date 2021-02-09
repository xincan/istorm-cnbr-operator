// Title  main.go
// Description  程序主入口
// Author  jiangxincan@hatech.com.cn  2021/1/22 11:36
// update  jiangxincan@hatech.com.cn  2021/1/22 11:36
package main

import (
	"github.com/sirupsen/logrus"
	_ "hatech.com.cn/istorm-cnbr-operator/docs"
	"hatech.com.cn/istorm-cnbr-operator/logs"
	"net/http"
)

// logrus提供了New()函数来创建一个logrus的实例。
// 项目中，可以创建任意数量的logrus实例。
var log = logrus.New()

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

	hook := logs.NewHook("d:/golog.log")
	logrus.WithField("file", "d:/log/golog.log").Info("日志信息存放地址: ")
	logrus.AddHook(hook)
	_ = http.ListenAndServe(":8080", Router())
}
