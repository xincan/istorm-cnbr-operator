// Title  nacosService.go
// Description  nacosService服务层
// Author  jiangxincan@hatech.com.cn  2021/1/22 11:36
// update  jiangxincan@hatech.com.cn  2021/1/22 11:36
package service

import (
	"github.com/sirupsen/logrus"
	_ "github.com/spf13/viper/remote"
	config "hatech.com.cn/istorm-cnbr-operator/config"
	"hatech.com.cn/istorm-cnbr-operator/vo"
)

// Title    		nacosService服务接口
// Description   	application/json编码解析函数
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
type INacosService interface {
	// 加法运算
	GetPerson(name string, age uint64) (*vo.NacosVo, error)
}

// Title    		nacosService服务实现结构体
// Description   	nacosService服务实现结构体
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
type NacosService struct {
}

func NacosServiceImpl() *NacosService {
	return &NacosService{}
}

// Title    			加法运算方法
// Description   		加法运算
// Auth      			jiangxincan@hatech.com.cn       			时间（2021/1/22 11:36）
// Return    			*vo.NacosVo     							"返回nacosVo结构体"
func (service *NacosService) GetPerson(name string, age uint64) (*vo.NacosVo, error) {
	person := vo.NacosVo{
		Name:       name,
		Age:        age,
		ConfigName: config.RemoteConfig.GetString("config.name"),
		ConfigAge:  config.RemoteConfig.GetUint64("config.age"),
	}
	logrus.WithField("data", person).Info("获取YAML配置文件参数")
	return &person, nil
}
