// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  xincan  2021/2/7 16:10
// @Update  xincan  2021/2/7 16:10
package goconvey

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/gavv/httpexpect/v2"
	_ "github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"hatech.com.cn/istorm-cnbr-operator/dto"
	"hatech.com.cn/istorm-cnbr-operator/kubernetes/transport"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

var number = &dto.KubernetesDto{
	Num1: 36,
	Num2: 4,
}

// mock handler 测试
func TestMockTransportFromKubernetesService(t *testing.T) {

	param := url.Values{}
	param.Set("num1", "36")
	param.Set("num2", "4")

	Convey("两个数处理，单元测试--transport", t, func() {

		Convey("将两数相加", func() {
			url, _ := url.ParseRequestURI("http://localhost:8080/add")
			url.RawQuery = param.Encode()
			urlStr := fmt.Sprintf("%v", url)
			req, _ := http.NewRequest("GET", urlStr, nil)
			req.Header.Add("Content-Type", "application/json")
			log.Print(urlStr)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			res, _ := transport.DecodeAddKubernetesRequest(ctx, req)
			r, _ := res.(dto.KubernetesDto)
			if r.Num1 == number.Num1 && r.Num2 == number.Num2 {
				log.Printf("测试成功：将两数相加。结构体为：%v", r)
			}
		})

		Convey("将两数相减", func() {
			url, _ := url.ParseRequestURI("http://localhost:8080/sub")
			url.RawQuery = param.Encode()
			urlStr := fmt.Sprintf("%v", url)
			req, _ := http.NewRequest("GET", urlStr, nil)
			req.Header.Add("Content-Type", "application/json")
			log.Print(urlStr)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			res, _ := transport.DecodeSubKubernetesRequest(ctx, req)
			r, _ := res.(dto.KubernetesDto)
			if r.Num1 == number.Num1 && r.Num2 == number.Num2 {
				log.Printf("测试成功：将两数相减。结构体为：%v", r)
			}
		})

		Convey("将两数相乘", func() {
			url, _ := url.ParseRequestURI("http://localhost:8080/mul")
			url.RawQuery = param.Encode()
			urlStr := fmt.Sprintf("%v", url)
			req, _ := http.NewRequest("POST", urlStr, nil)
			req.Header.Add("Content-Type", "application/json")
			log.Print(urlStr)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			res, _ := transport.DecodeMulKubernetesRequest(ctx, req)
			r, _ := res.(dto.KubernetesDto)
			if r.Num1 == number.Num1 && r.Num2 == number.Num2 {
				log.Printf("测试成功：将两数相乘。结构体为：%v", r)
			}
		})

		Convey("将两数相除", func() {
			url, _ := url.ParseRequestURI("http://localhost:8080/div")
			urlStr := fmt.Sprintf("%v", url)
			a, _ := json.Marshal(number)
			body := strings.TrimSpace(string(a))
			req, _ := http.NewRequest("POST", urlStr, strings.NewReader(body))
			req.Header.Add("Content-Type", "application/json")
			log.Print(urlStr)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			res, _ := transport.DecodeDivKubernetesRequest(ctx, req)
			r, _ := res.(dto.KubernetesDto)
			if r.Num1 == number.Num1 && r.Num2 == number.Num2 {
				log.Printf("测试成功：将两数相除。结构体为：%v", r)
			}

		})

		Convey("求两个数的平均值", func() {
			url, _ := url.ParseRequestURI("http://localhost:8080/ave")
			url.RawQuery = param.Encode()
			urlStr := fmt.Sprintf("%v", url)
			req, _ := http.NewRequest("POST", urlStr, nil)
			req.Header.Add("Content-Type", "application/json")
			log.Print(urlStr)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			res, _ := transport.DecodeAveKubernetesRequest(ctx, req)
			r, _ := res.(dto.KubernetesDto)
			if r.Num1 == number.Num1 && r.Num2 == number.Num2 {
				log.Printf("测试成功：求两个数的平均值。结构体为：%v", r)
			}
		})

	})

}
