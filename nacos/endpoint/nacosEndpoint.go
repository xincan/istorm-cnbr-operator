// Title  nacosEndpoint.go
// Description  nacosEndpoint端点
// Author  jiangxincan@hatech.com.cn  2022/02/17 11:36
// update  jiangxincan@hatech.com.cn  2022/02/17 11:36
package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"xincan.com.cn/istorm-cnbr-operator/dto"
	"xincan.com.cn/istorm-cnbr-operator/nacos/service"
	resp "xincan.com.cn/istorm-cnbr-operator/utils/response"
)

// Title    		获取Nacos远程人员配置信息函数
// Description   	获取人员信息
// Auth      		jiangxincan@hatech.com.cn         	时间（2022/02/17 11:36）
// Param    		nacosService     					service.InacosService    	"获取Nacos远程配置信息人员接口"
// Return     		responseResult        				resp.ResponseResult       	"返回统一封装结果"
func GetPersonEndpoint(nacosService *service.NacosService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(dto.NacosDto)
		person, _ := nacosService.GetPerson(req.Name, req.Age)
		return resp.ResultSuccess(person), nil
	}
}
