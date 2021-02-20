// Title  router.go
// Description  路由配置
// Author  jiangxincan@hatech.com.cn  2021/1/22 11:36
// update  jiangxincan@hatech.com.cn  2021/1/22 11:36
package main

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"hatech.com.cn/istorm-cnbr-operator/kubernetes/handler"
)

// Title    		路由列表
// Description   	项目路由列表配置
// Auth      		jiangxincan@hatech.com.cn         时间（2021/1/22 11:36）
// Return    		router        	*mux.Router     "返回路由对象"
func Router() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DomID("#swagger-ui"),
	))
	// kubernetes路由列表
	router.Methods("GET").Path("/add").Handler(handler.AddKubernetesHandler())
	router.Methods("GET").Path("/sub").Handler(handler.SubKubernetesHandler())
	router.Methods("POST").Path("/mul").Handler(handler.MulKubernetesHandler())
	router.Methods("POST").Path("/div").Handler(handler.DivKubernetesHandler())
	router.Methods("POST").Path("/ave").Handler(handler.AveKubernetesHandler())
	return router
}
