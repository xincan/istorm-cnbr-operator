// Title  nacosHandler.go
// Description  nacosHandler协调适配器
// Author  jiangxincan@hatech.com.cn  2021/1/22 11:36
// update  jiangxincan@hatech.com.cn  2021/1/22 11:36
package handler

import (
	"github.com/go-kit/kit/transport/http"
	"hatech.com.cn/istorm-cnbr-operator/nacos/endpoint"
	"hatech.com.cn/istorm-cnbr-operator/nacos/service"
	"hatech.com.cn/istorm-cnbr-operator/nacos/transport"
	"hatech.com.cn/istorm-cnbr-operator/utils"
)

// Title    		求和协调适配器函数
// Description   	求和协调适配器函数
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// Return    		httpServer        *http.Server    "求和服务"
func GetPersonHandler() *http.Server {
	nacosService := service.NacosServiceImpl()
	add := endpoint.GetPersonEndpoint(nacosService)
	return http.NewServer(add, transport.GetPersonRequest, utils.EncodeNacosResponse)
}
