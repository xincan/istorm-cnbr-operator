// Title  kubernetesEndpoint.go
// Description  KubernetesEndpoint端点
// Author  jiangxincan@hatech.com.cn  2021/1/22 11:36
// update  jiangxincan@hatech.com.cn  2021/1/22 11:36
package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"hatech.com.cn/istorm-cnbr-operator/dto"
	"hatech.com.cn/istorm-cnbr-operator/kubernetes/service"
	resp "hatech.com.cn/istorm-cnbr-operator/utils/response"
)

// Title    		求和结果端点函数
// Description   	两个数相加求和
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// Param    		kubernetesService     service.IKubernetesService    "计算机接口"
// Return     		responseResult        resp.ResponseResult       	"返回统一封装结果"
func MakeAddKubernetesEndpoint(kubernetesService service.IKubernetesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.KubernetesDto)
		result, _ := kubernetesService.Add(req.Num1, req.Num2)
		return resp.ResultSuccess(result), nil
	}
}

// Title    		求差结果端点函数
// Description   	两个数相减求差
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// Param    		kubernetesService     service.IKubernetesService    "计算机接口"
// Return     		responseResult        resp.ResponseResult       	"返回统一封装结果"
func MakeSubKubernetesEndpoint(kubernetesService service.IKubernetesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.KubernetesDto)
		result, _ := kubernetesService.Sub(req.Num1, req.Num2)
		return resp.ResultSuccess(result), nil
	}
}

// Title    		求乘积结果端点函数
// Description   	两个数相乘求乘积
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// Param    		kubernetesService     service.IKubernetesService    "计算机接口"
// Return     		responseResult        resp.ResponseResult       	"返回统一封装结果"
func MakeMulKubernetesEndpoint(kubernetesService service.IKubernetesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.KubernetesDto)
		result, _ := kubernetesService.Mul(req.Num1, req.Num2)
		return resp.ResultSuccess(result), nil
	}
}

// Title    		求商结果端点函数
// Description   	两个数相除求商
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// Param    		kubernetesService     service.IKubernetesService    "计算机接口"
// Return     		responseResult        resp.ResponseResult	        "返回统一封装结果"
func MakeDivKubernetesEndpoint(kubernetesService service.IKubernetesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.KubernetesDto)
		result, e := kubernetesService.Div(req.Num1, req.Num2)
		if result == nil && e != nil {
			return resp.ResultError(e, req), e
		}
		return resp.ResultSuccess(result), nil
	}
}

// title    		求平均值结果端点函数
// description   	两个数相加求和之后在求平均值
// auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// param    		kubernetesService     service.IKubernetesService    "计算机接口"
// return     		responseResult        resp.ResponseResult	        "返回统一封装结果"
func MakeAveKubernetesEndpoint(kubernetesService service.IKubernetesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.KubernetesDto)
		result, _ := kubernetesService.Ave(req.Num1, req.Num2)
		return resp.ResultSuccess(result), nil
	}
}
