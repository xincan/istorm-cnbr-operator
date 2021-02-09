// Title  kubernetesService.go
// Description  kubernetesService服务层
// Author  jiangxincan@hatech.com.cn  2021/1/22 11:36
// update  jiangxincan@hatech.com.cn  2021/1/22 11:36
package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"hatech.com.cn/istorm-cnbr-operator/vo"
	"log"
)

// Title    		kubernetesService服务接口
// Description   	application/json编码解析函数
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
type IKubernetesService interface {
	// 加法运算
	Add(num1 int, num2 int) (*vo.KubernetesVo, error)
	// 减法运算
	Sub(num1 int, num2 int) (*vo.KubernetesVo, error)
	// 乘法运算
	Mul(num1 int, num2 int) (*vo.KubernetesVo, error)
	// 除法运算
	Div(num1 int, num2 int) (*vo.KubernetesVo, error)
	// 平均值运算
	Ave(num1 int, num2 int) (*vo.KubernetesVo, error)
}

// Title    		kubernetesService服务实现结构体
// Description   	kubernetesService服务实现结构体
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
type kubernetesService struct {
}

func KubernetesServiceImpl() *kubernetesService {
	return &kubernetesService{}
}

// Title    			加法运算方法
// Description   		加法运算
// Auth      			jiangxincan@hatech.com.cn       			时间（2021/1/22 11:36）
// Return    			kubernetesVo        	*vo.KubernetesVo     "返回kubernetesVo结构体"
func (service *kubernetesService) Add(num1 int, num2 int) (*vo.KubernetesVo, error) {
	log.Printf("service: num1: %v, num2: %v", num1, num2)
	return &vo.KubernetesVo{Result: num1 + num2}, nil
}

// Title    			减法运算方法
// Description   		减法运算
// Auth      			jiangxincan@hatech.com.cn       			时间（2021/1/22 11:36）
// Return    			kubernetesVo        	*vo.KubernetesVo     "返回kubernetesVo结构体"
func (service *kubernetesService) Sub(num1 int, num2 int) (*vo.KubernetesVo, error) {
	return &vo.KubernetesVo{Result: num1 - num2}, nil
}

// Title    			乘法运算方法
// Description   		乘法运算
// Auth      			jiangxincan@hatech.com.cn       			时间（2021/1/22 11:36）
// Return    			kubernetesVo        	*vo.KubernetesVo     "返回kubernetesVo结构体"
func (service *kubernetesService) Mul(num1 int, num2 int) (*vo.KubernetesVo, error) {
	return &vo.KubernetesVo{Result: num1 * num2}, nil
}

// Title    			除法法运算方法
// Description   		除法运算
// Auth      			jiangxincan@hatech.com.cn       			时间（2021/1/22 11:36）
// Return    			kubernetesVo        	*vo.KubernetesVo     "返回kubernetesVo结构体"
func (service *kubernetesService) Div(num1 int, num2 int) (v *vo.KubernetesVo, err error) {
	logrus.WithFields(logrus.Fields{"num1": num1, "num2": num2}).Info("求两个数的商", ": ", "除数不能为0")
	if num2 == 0 {
		return nil, errors.New("除数不能为0")
	}
	return &vo.KubernetesVo{Result: num1 / num2}, nil
}

// Title    			平均值运算方法
// Description   		平均值运算
// Auth      			jiangxincan@hatech.com.cn       			时间（2021/1/22 11:36）
// Return    			kubernetesVo        	*vo.KubernetesVo     "返回kubernetesVo结构体"
func (service *kubernetesService) Ave(num1 int, num2 int) (*vo.KubernetesVo, error) {
	return &vo.KubernetesVo{Result: (num1 + num2) / 2}, nil
}
