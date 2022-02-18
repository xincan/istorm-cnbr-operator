// Title  nacosTransport.go
// Description  传输层
// Author  jiangxincan@hatech.com.cn  2022/02/17 11:36
// update  jiangxincan@hatech.com.cn  2022/02/17 11:36
package transport

import (
	"context"
	"github.com/sirupsen/logrus"
	"hatech.com.cn/istorm-cnbr-operator/dto"
	"net/http"
	"strconv"
)

// @Title GetPersonRequest 获取人员信息
// @Summary 获取人员信息(get请求，【query，application/json方式】，路径中有参数)
// @Description 获取人员信息
// @Tags Nacos获取人员信息
// @Accept application/json
// @Produce application/json
// @Param name query string true "姓名" default(张三)
// @Param age query uint64 true "年龄" default(44)
// @Success 200 {interface} response.ResponseResult
// @Router /person [get]
// @Auth      		jiangxincan@hatech.com.cn         	时间（2022/02/17 11:36）
// @Return			dto.NacosDto				  		"person传参对象"
func GetPersonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	name := r.FormValue("name")
	age, _ := strconv.ParseUint(r.FormValue("age"), 10, 64)
	logrus.WithFields(logrus.Fields{"name": name, "age": age}).Info("获取前端传参")
	return dto.NacosDto{Name: name, Age: age}, nil
}
