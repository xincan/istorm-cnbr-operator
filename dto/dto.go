// Title  vo.go
// Description  vo传参包装包
// Author  jiangxincan@hatech.com.cn  2021/1/22 11:36
// update  jiangxincan@hatech.com.cn  2021/1/22 11:36
package dto

// Title    		kubernetes传参结构体
// Description   	kubernetes传参结构体
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
type KubernetesDto struct {
	Num1 int `json:"num1" example:"36"` // 第一个殊字
	Num2 int `json:"num2" example:"4"`  // 第二个殊字
}

// Title    		nacos传参结构体
// Description   	nacos传参结构体
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
type NacosDto struct {
	Name string `json:"name" example:"zhangsan"`
	Age  uint64 `json:"age" example:36`
}
