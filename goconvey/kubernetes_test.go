// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  xincan  2021/2/7 16:10
// @Update  xincan  2021/2/7 16:10
package goconvey

import (
	"errors"
	_ "github.com/gavv/httpexpect/v2"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"hatech.com.cn/istorm-cnbr-operator/kubernetes/service"
	"hatech.com.cn/istorm-cnbr-operator/vo"
	"testing"
)

// mock测试
func TestMockDivFromKubernetesService(t *testing.T) {

	ctrl := gomock.NewController(t)
	m := NewMockIKubernetesService(ctrl)

	Convey("两个数处理", t, func() {

		Convey("将两数相加", func() {
			m.EXPECT().Add(1, 2).Return(&vo.KubernetesVo{Result: 3}, nil)
			res, _ := service.KubernetesServiceImpl().Add(1, 2)
			log.Print(res.Result)

			if res.Result != 3 {
				t.Error("连个数相加，结果为：%i", res.Result)
			}
		})

		Convey("将两数相减", func() {
			m.EXPECT().Sub(10, 2).Return(&vo.KubernetesVo{Result: 8}, nil)
			res, _ := service.KubernetesServiceImpl().Sub(10, 2)
			log.Print(res.Result)
			if res.Result != 8 {
				t.Error("连个数相减，结果为：%i", res.Result)
			}
		})

		Convey("将两数相乘", func() {
			m.EXPECT().Mul(10, 2).Return(&vo.KubernetesVo{Result: 20}, nil)
			res, _ := service.KubernetesServiceImpl().Mul(10, 2)
			log.Print(res.Result)
			if res.Result != 20 {
				t.Error("连个数相乘，结果为：%i", res.Result)
			}
		})

		Convey("将两数相除", func() {

			Convey("将两数相除，除数不为：0", func() {
				m.EXPECT().Div(10, 2).Return(&vo.KubernetesVo{Result: 5}, nil)
				res, _ := service.KubernetesServiceImpl().Div(10, 2)
				log.Print(res.Result)
				if res.Result != 5 {
					t.Error("连个数相除，结果为：%i", res.Result)
				}
			})

			Convey("将两数相除，除数是为：0", func() {
				m.EXPECT().Div(10, 0).Return(&vo.KubernetesVo{Result: 0}, errors.New("除数不能为0"))
				_, err := service.KubernetesServiceImpl().Div(10, 0)
				log.Print(err.Error())
				if err.Error() == "除数不能为0" {
					log.Print("两个数相除，除数不能为：0，目前除数为：", 0)
				}
			})

		})

		Convey("求两个数的平均值", func() {
			m.EXPECT().Ave(10, 2).Return(&vo.KubernetesVo{Result: 6}, nil)
			res, _ := service.KubernetesServiceImpl().Ave(10, 2)
			log.Print(res.Result)
			if res.Result != 6 {
				t.Error("求两个数的平均值，结果为：%i", res.Result)
			}
		})

	})

}
