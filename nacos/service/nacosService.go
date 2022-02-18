// Title  nacosService.go
// Description  nacosService服务层
// Author  jiangxincan@hatech.com.cn  2022/02/17 11:36
// update  jiangxincan@hatech.com.cn  2022/02/17 11:36
package service

import (
	_ "github.com/spf13/viper/remote"
	"hatech.com.cn/istorm-cnbr-operator/config"
	"hatech.com.cn/istorm-cnbr-operator/vo"
)

// Title    		nacosService服务接口
// Description   	application/json编码解析函数
// Auth      		jiangxincan@hatech.com.cn         时间（2022/02/17 11:36）
type INacosService interface {
	// 获取人员信息
	GetPerson(name string, age uint64) (*vo.NacosVo, error)
}

// Title    		nacosService服务实现结构体
// Description   	nacosService服务实现结构体
// Auth      		jiangxincan@hatech.com.cn         时间（2022/02/17 11:36）
type NacosService struct{}

func NacosServiceImpl() *NacosService {
	return &NacosService{}
}

// Title    			获取远程Nacos配置信息
// Description   		获取配置信息
// Auth      			jiangxincan@hatech.com.cn       			时间（2022/02/17 11:36）
// Return    			*vo.NacosVo     							"返回nacosVo结构体"
func (service *NacosService) GetPerson(name string, age uint64) (*vo.NacosVo, error) {
	person := vo.NacosVo{
		Name:       name,
		Age:        age,
		ConfigName: config.RemoteConfig.GetString("config.name"),
		ConfigAge:  config.RemoteConfig.GetUint64("config.age"),
	}
	return &person, nil
}
