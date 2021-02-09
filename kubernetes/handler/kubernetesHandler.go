// Title  kubernetesHandler.go
// Description  KubernetesHandler协调适配器
// Author  jiangxincan@hatech.com.cn  2021/1/22 11:36
// update  jiangxincan@hatech.com.cn  2021/1/22 11:36
package handler

import (
	"github.com/go-kit/kit/transport/http"
	"hatech.com.cn/istorm-cnbr-operator/kubernetes/endpoint"
	"hatech.com.cn/istorm-cnbr-operator/kubernetes/service"
	"hatech.com.cn/istorm-cnbr-operator/kubernetes/transport"
	"hatech.com.cn/istorm-cnbr-operator/utils"
)

// Title    		求和协调适配器函数
// Description   	求和协调适配器函数
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// Return    		httpServer        *http.Server    "求和服务"
func AddKubernetesHandler() *http.Server {
	kubernetesService := service.KubernetesServiceImpl()
	add := endpoint.MakeAddKubernetesEndpoint(kubernetesService)
	return http.NewServer(add, transport.DecodeAddKubernetesRequest, utils.EncodeKubernetesResponse)
}

// Title    		求差协调适配器函数
// Description   	求差协调适配器函数
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// Return    		httpServer        *http.Server    "求差服务"
func SubKubernetesHandler() *http.Server {
	kubernetesService := service.KubernetesServiceImpl()
	sub := endpoint.MakeSubKubernetesEndpoint(kubernetesService)
	return http.NewServer(sub, transport.DecodeSubKubernetesRequest, utils.EncodeKubernetesResponse)
}

// Title    		求乘积协调适配器函数
// Description   	求乘积协调适配器函数
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// Return    		httpServer        *http.Server    "乘积服务"
func MulKubernetesHandler() *http.Server {
	kubernetesService := service.KubernetesServiceImpl()
	mul := endpoint.MakeMulKubernetesEndpoint(kubernetesService)
	return http.NewServer(mul, transport.DecodeMulKubernetesRequest, utils.EncodeKubernetesResponse)
}

// Title    		求商协调适配器函数
// Description   	求商协调适配器函数
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// Return    		httpServer        *http.Server    "求商服务"
func DivKubernetesHandler() *http.Server {
	kubernetesService := service.KubernetesServiceImpl()
	div := endpoint.MakeDivKubernetesEndpoint(kubernetesService)
	return http.NewServer(div, transport.DecodeDivKubernetesRequest, utils.EncodeKubernetesResponse)
}

// Title    		求平均值协调适配器函数
// Description   	求平均值协调适配器函数
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// Return    		httpServer        *http.Server    "求商服务"
func AveKubernetesHandler() *http.Server {
	kubernetesService := service.KubernetesServiceImpl()
	div := endpoint.MakeAveKubernetesEndpoint(kubernetesService)
	return http.NewServer(div, transport.DecodeAveKubernetesRequest, utils.EncodeKubernetesResponse)
}
