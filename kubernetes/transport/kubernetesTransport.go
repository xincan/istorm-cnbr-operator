// Title  kubernetesTransport.go
// Description  传输层
// Author  jiangxincan@hatech.com.cn  2021/1/22 11:36
// update  jiangxincan@hatech.com.cn  2021/1/22 11:36
package transport

import (
	"context"
	"encoding/json"
	"hatech.com.cn/istorm-cnbr-operator/dto"
	"io/ioutil"
	"net/http"
	"strconv"
)

// @Title DecodeAddKubernetesRequest 求和函数
// @Summary 求和函数(get请求，【query，application/json方式】，路径中有参数)
// @Description 根据两个数字求和
// @Tags 运算相关函数
// @Accept application/json
// @Produce application/json
// @Param num1 query int true "第一个数字" default(36)
// @Param num2 query int true "第二个数字" default(4)
// @Success 200 {interface} response.ResponseResult
// @Router /add [get]
// @Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// @Return			dto.KubernetesDto				  "kubernetes传参对象"
func DecodeAddKubernetesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	num1, _ := strconv.Atoi(r.FormValue("num1"))
	num2, _ := strconv.Atoi(r.FormValue("num2"))
	return dto.KubernetesDto{Num1: num1, Num2: num2}, nil
}

// @Title DecodeSubKubernetesRequest 求差函数
// @Summary 求差函数(get请求，【query，application/json方式】，路径上没有参数)
// @Description 根据两个殊字求两个数字之减
// @Tags 运算相关函数
// @Accept application/json
// @Produce application/json
// @Param num1 query int true "第一个数字" default(36)
// @Param num2 query int true "第二个数字" default(4)
// @Success 200 {interface} response.ResponseResult
// @Router /sub [get]
// @Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// @Return			dto.KubernetesDto				  "kubernetes传参对象"
func DecodeSubKubernetesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	num1, _ := strconv.Atoi(r.FormValue("num1"))
	num2, _ := strconv.Atoi(r.FormValue("num2"))
	return dto.KubernetesDto{Num1: num1, Num2: num2}, nil
}

// @Title DecodeMulKubernetesRequest 乘积函数
// @Summary 求乘积函数(post请求，【query，application/json方式】，路径有参数)
// @Description 根据两个数字求乘积
// @Tags 运算相关函数
// @Accept application/json
// @Produce application/json
// @Param num1 query int true "第一个数字" default(36)
// @Param num2 query int true "第二个数字" default(4)
// @Success 200 {interface} response.ResponseResult
// @Router /mul [post]
// @Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// @Return			dto.KubernetesDto				  "kubernetes传参对象"
func DecodeMulKubernetesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	num1, _ := strconv.Atoi(r.FormValue("num1"))
	num2, _ := strconv.Atoi(r.FormValue("num2"))
	return dto.KubernetesDto{Num1: num1, Num2: num2}, nil
}

// @Title DecodeDivKubernetesRequest 除法函数
// @Summary 求商函数(post请求，【body, application/json方式】，路径没有参数)
// @Description 根据两个数字求商
// @Tags 运算相关函数
// @Accept application/json
// @Produce application/json
// @Param kubernetes body dto.KubernetesDto true "k8s对象"
// @Success 200 {interface} response.ResponseResult
// @Router /div [post]
// @Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// @Return			dto.KubernetesDto				  "kubernetes传参对象"
func DecodeDivKubernetesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body, _ := ioutil.ReadAll(r.Body)
	var kubernetesDto dto.KubernetesDto
	_ = json.Unmarshal(body, &kubernetesDto)
	return kubernetesDto, nil
}

// @Title DecodeDivKubernetesRequest 平均值函数
// @Summary 求平均值函数(post请求，【query，application/x-www-form-urlencoded方式】，路径没有参数)
// @Description 根据两个数字求平均值
// @Tags 运算相关函数
// @Accept application/x-www-form-urlencoded
// @Produce application/x-www-form-urlencoded
// @Param num1 query int true "第一个数字" default(36)
// @Param num2 query int true "第二个数字" default(4)
// @Success 200 {interface} response.ResponseResult
// @Router /ave [post]
// @Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// @Return			dto.KubernetesDto				  "kubernetes传参对象"
func DecodeAveKubernetesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	num1, _ := strconv.Atoi(r.FormValue("num1"))
	num2, _ := strconv.Atoi(r.FormValue("num2"))
	return dto.KubernetesDto{Num1: num1, Num2: num2}, nil
}
