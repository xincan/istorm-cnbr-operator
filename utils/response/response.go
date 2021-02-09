// Title  response.go
// Description  统一返回值包
// Author  jiangxincan@hatech.com.cn  2021/1/22 11:36
// update  jiangxincan@hatech.com.cn  2021/1/22 11:36
package response

import (
	json "encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
)

const (
	requestSuccess = 200 // 成功返回状态码
	requestError   = 500 // 失败返回状态码

	operateSuccessMessage = "操作成功" // 成功返回信息
	operateErrorMessage   = "操作失败" // 失败返回信息
)

// Title    		统一返回结果结构体
// Description   	统一返回结果结构体
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
type result struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"操作成功"`
	Count   int         `json:"count" example:"1"`
	Data    interface{} `json:"data" example:"nil"`
}

// Title    		编码解析函数
// Description   	application/json编码解析函数
// Auth      		jiangxincan@hatech.com.cn         	时间（2021/1/22 11:36）
// Param    		err     				error    	"异常对象"
// Return    		message        			string     	"返回消息说明"
func message(err error) string {
	if err != nil {
		return operateErrorMessage + ": " + err.Error()
	}
	return operateSuccessMessage
}

// Title    		统一错误信息日志记录函数
// Description   	统一错误信息日志记录函数
// Auth      		jiangxincan@hatech.com.cn         			时间（2021/1/22 11:36）
// Param    		message     			string    			"异常对象"
// Param    		data        			interface{}     	"结果对象"
func logError(msg string, data interface{}) *result {
	var mapResult map[string]interface{}
	err := errors.New(msg)
	da, _ := json.Marshal(data)
	_ = json.Unmarshal(da, &mapResult)
	log.WithFields(mapResult).Error("出现异常：", err)
	return &result{requestError, message(err), 0, data}
}

// Title    		统一返回错误信息日志记录函数
// Description   	统一返回错误信息日志记录函数
// Auth      		jiangxincan@hatech.com.cn         			时间（2021/1/22 11:36）
// Param    		err     				error    			"异常对象"
// Param    		data        			interface{}     	"结果对象"
func ResultError(err error, data interface{}) *result {
	return logError(err.Error(), data)
}

// Title    		统一返回成功信息日志记录函数
// Description   	统一返回成功信息日志记录函数
// Auth      		jiangxincan@hatech.com.cn         			时间（2021/1/22 11:36）
// Param    		data        			interface{}     	"结果对象"
func ResultSuccess(data interface{}) *result {
	return &result{requestSuccess, message(nil), 1, data}
}

// Title    		统一返回成功信息日志记录函数
// Description   	统一返回成功信息日志记录函数
// Auth      		jiangxincan@hatech.com.cn         		时间（2021/1/22 11:36）
// Param    		count     				int    			"异常对象"
// Param    		data        			interface{}     "结果对象"
func ResultSuccessManyAndCount(count int, data interface{}) *result {
	return &result{requestSuccess, message(nil), count, data}
}
