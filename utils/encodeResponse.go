// Title  encodeResponse.go
// Description  编码解析
// Author  jiangxincan@hatech.com.cn  2021/1/22 11:36
// update  jiangxincan@hatech.com.cn  2021/1/22 11:36
package utils

import (
	"context"
	"encoding/json"
	"net/http"
)

// Title    		编码解析函数
// Description   	application/json编码解析函数
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// Return    		error
func EncodeKubernetesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

// Title    		编码解析函数
// Description   	application/json编码解析函数
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// Return    		error
func EncodeNacosResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
