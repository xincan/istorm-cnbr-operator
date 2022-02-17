// Title  vo.go
// Description  vo视图包
// Author  jiangxincan@hatech.com.cn  2021/1/22 11:36
// update  jiangxincan@hatech.com.cn  2021/1/22 11:36
package vo

// Title    		kubernetes视图结构体
// Description   	kubernetes视图结构体
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
type KubernetesVo struct {
	Result int `json:"result"` // 返回结果值
}

// Title    		Nacos视图结构体
// Description   	Nacos视图结构体
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
type NacosVo struct {
	Name       string `json:"name" example:"zhangsan"`
	Age        uint64 `json:"age" example:36`
	ConfigName string `json:"configName" example:"zhangsan"`
	ConfigAge  uint64 `json:"configAge" example:36`
}
