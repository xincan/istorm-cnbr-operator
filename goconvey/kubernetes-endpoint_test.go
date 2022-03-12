// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  xincan  2021/2/7 16:10
// @Update  xincan  2021/2/7 16:10
package goconvey

import (
	"context"
	_ "github.com/gavv/httpexpect/v2"
	_ "github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
	"xincan.com.cn/istorm-cnbr-operator/dto"
	ep "xincan.com.cn/istorm-cnbr-operator/kubernetes/endpoint"
	"xincan.com.cn/istorm-cnbr-operator/kubernetes/service"
	resp "xincan.com.cn/istorm-cnbr-operator/utils/response"
	"xincan.com.cn/istorm-cnbr-operator/vo"
)

var kdt = dto.KubernetesDto{
	Num1: 36,
	Num2: 4,
}

// mock handler 测试
func TestMockEndpointFromKubernetesService(t *testing.T) {

	Convey("两个数处理,单元测试--endpoint", t, func() {

		Convey("将两数相加", func() {
			svc := service.KubernetesServiceImpl()
			endpoint := ep.MakeAddKubernetesEndpoint(svc)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			data, _ := endpoint(ctx, kdt)
			resVo := data.(*resp.Result)
			resData := (resVo.Data).(*vo.KubernetesVo)
			if resVo.Code == 200 && resData.Result == 40 {
				log.Printf("测试成功：将两数相加。结果为： %v + %v = %v", kdt.Num1, kdt.Num2, resData.Result)
			}
		})

		Convey("将两数相减", func() {
			svc := service.KubernetesServiceImpl()
			endpoint := ep.MakeSubKubernetesEndpoint(svc)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			data, _ := endpoint(ctx, kdt)
			resVo := data.(*resp.Result)
			resData := (resVo.Data).(*vo.KubernetesVo)
			if resVo.Code == 200 && resData.Result == 32 {
				log.Printf("测试成功：将两数相减。结果为： %v - %v = %v", kdt.Num1, kdt.Num2, resData.Result)
			}
		})

		Convey("将两数相乘", func() {
			svc := service.KubernetesServiceImpl()
			endpoint := ep.MakeMulKubernetesEndpoint(svc)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			data, _ := endpoint(ctx, kdt)
			resVo := data.(*resp.Result)
			resData := (resVo.Data).(*vo.KubernetesVo)
			if resVo.Code == 200 && resData.Result == 144 {
				log.Printf("测试成功：将两数相乘。结果为： %v * %v = %v", kdt.Num1, kdt.Num2, resData.Result)
			}
		})

		Convey("将两数相除", func() {

			Convey("将两数相除，除数不为：0", func() {
				svc := service.KubernetesServiceImpl()
				endpoint := ep.MakeDivKubernetesEndpoint(svc)
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				data, _ := endpoint(ctx, kdt)
				resVo := data.(*resp.Result)
				resData := (resVo.Data).(*vo.KubernetesVo)
				if resVo.Code == 200 && resData.Result == 9 {
					log.Printf("测试成功：将两数相除，除数不为：0。结果为： %v ➗ %v = %v", kdt.Num1, kdt.Num2, resData.Result)
				}
			})

			Convey("将两数相除，除数为：0", func() {
				kdt.Num2 = 0
				svc := service.KubernetesServiceImpl()
				endpoint := ep.MakeDivKubernetesEndpoint(svc)
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				data, err := endpoint(ctx, kdt)
				resVo := data.(*resp.Result)
				if resVo.Code == 500 {
					log.Printf("测试成功：异常测试，将两数相除，除数为：0。 异常信息：%v。 Num1 = %v, Num2 = %v", err, kdt.Num1, kdt.Num2)
				}
			})

		})

		Convey("求两个数的平均值", func() {
			kdt.Num2 = 4
			svc := service.KubernetesServiceImpl()
			endpoint := ep.MakeAveKubernetesEndpoint(svc)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			data, _ := endpoint(ctx, kdt)
			resVo := data.(*resp.Result)
			resData := (resVo.Data).(*vo.KubernetesVo)
			if resVo.Code == 200 && resData.Result == 20 {
				log.Printf("测试成功：求两个数的平均值。结果为： (%v + %v) / 2 = %v", kdt.Num1, kdt.Num2, resData.Result)
			}
		})

	})

}
