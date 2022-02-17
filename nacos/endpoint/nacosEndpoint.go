// Title  nacosEndpoint.go
// Description  nacosEndpoint端点
// Author  jiangxincan@hatech.com.cn  2021/1/22 11:36
// update  jiangxincan@hatech.com.cn  2021/1/22 11:36
package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"hatech.com.cn/istorm-cnbr-operator/dto"
	"hatech.com.cn/istorm-cnbr-operator/nacos/service"
	resp "hatech.com.cn/istorm-cnbr-operator/utils/response"
)

// Title    		求和结果端点函数
// Description   	两个数相加求和
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// Param    		nacosService     service.InacosService    "计算机接口"
// Return     		responseResult        resp.ResponseResult       	"返回统一封装结果"
func GetPersonEndpoint(nacosService *service.NacosService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(dto.NacosDto)
		person, _ := nacosService.GetPerson(req.Name, req.Age)
		return resp.ResultSuccess(person), nil
	}
}
