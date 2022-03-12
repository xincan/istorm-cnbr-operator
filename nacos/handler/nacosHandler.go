// Title  nacosHandler.go
// Description  nacosHandler协调适配器
// Author  jiangxincan@hatech.com.cn  2022/02/17 11:36
// update  jiangxincan@hatech.com.cn  2022/02/17 11:36
package handler

import (
	"github.com/go-kit/kit/transport/http"
	"xincan.com.cn/istorm-cnbr-operator/nacos/endpoint"
	"xincan.com.cn/istorm-cnbr-operator/nacos/service"
	"xincan.com.cn/istorm-cnbr-operator/nacos/transport"
	"xincan.com.cn/istorm-cnbr-operator/utils"
)

// Title    		获取人员信息协调适配器函数
// Description   	获取人员信息协调适配器函数
// Auth      		jiangxincan@hatech.com.cn         	时间（2022/02/17 11:36）
// Return    		httpServer        					*http.Server    "求和服务"
func GetPersonHandler() *http.Server {
	nacosService := service.NacosServiceImpl()
	add := endpoint.GetPersonEndpoint(nacosService)
	return http.NewServer(add, transport.GetPersonRequest, utils.EncodeNacosResponse)
}
