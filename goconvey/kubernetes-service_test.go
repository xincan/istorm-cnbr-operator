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
	"hatech.com.cn/istorm-cnbr-operator/dto"
	"hatech.com.cn/istorm-cnbr-operator/vo"
	"testing"
)

var param = dto.KubernetesDto{
	Num1: 36,
	Num2: 4,
}

var rest = struct {
	addRes     int
	subRes     int
	mulRes     int
	divHaveRes int
	divEqRes   int
	aveRes     int
}{
	addRes:     40,
	subRes:     32,
	mulRes:     144,
	divHaveRes: 9,
	divEqRes:   0,
	aveRes:     20,
}

// mock service 测试
func TestMockServiceKubernetesService(t *testing.T) {

	ctrl := gomock.NewController(t)
	mock := NewMockIKubernetesService(ctrl)

	Convey("两个数处理,单元测试--service", t, func() {

		Convey("将两数相加", func() {
			mock.EXPECT().Add(param.Num1, param.Num2).Return(&vo.KubernetesVo{Result: rest.addRes}, nil)
			data, _ := mock.Add(param.Num1, param.Num2)
			if data.Result == rest.addRes {
				log.Printf("测试成功：两个数相加。结果为：%v + %v = %v", param.Num1, param.Num2, data.Result)
			}
		})

		Convey("将两数相减", func() {
			mock.EXPECT().Sub(param.Num1, param.Num2).Return(&vo.KubernetesVo{Result: rest.subRes}, nil)
			data, _ := mock.Sub(param.Num1, param.Num2)
			if data.Result == rest.subRes {
				log.Printf("测试成功：两个数相减。结果为：%v - %v = %v", param.Num1, param.Num2, data.Result)
			}
		})

		Convey("将两数相乘", func() {
			mock.EXPECT().Mul(param.Num1, param.Num2).Return(&vo.KubernetesVo{Result: rest.mulRes}, nil)
			data, _ := mock.Mul(param.Num1, param.Num2)
			if data.Result == rest.mulRes {
				log.Printf("测试成功：两个数相乘。结果为：%v * %v = %v", param.Num1, param.Num2, data.Result)
			}
		})

		Convey("将两数相除", func() {

			Convey("将两数相除，除数不为：0", func() {
				mock.EXPECT().Div(param.Num1, param.Num2).Return(&vo.KubernetesVo{Result: rest.divHaveRes}, nil)
				data, _ := mock.Div(param.Num1, param.Num2)
				if data.Result == rest.divHaveRes {
					log.Printf("测试成功：两个数相除，除数不为：0。结果为：%v ➗ %v = %v", param.Num1, param.Num2, data.Result)
				}
			})

			Convey("将两数相除，除数是为：0", func() {
				mock.EXPECT().Div(param.Num1, 0).Return(&vo.KubernetesVo{Result: 0}, errors.New("除数不能为0"))
				_, err := mock.Div(param.Num1, 0)
				if err.Error() == "除数不能为0" {
					log.Print("测试成功：两个数相除，除数为：0。目前除数为：", 0)
				}
			})

		})

		Convey("求两个数的平均值", func() {
			mock.EXPECT().Ave(param.Num1, param.Num2).Return(&vo.KubernetesVo{Result: rest.aveRes}, nil)
			data, _ := mock.Ave(param.Num1, param.Num2)
			if data.Result == rest.aveRes {
				log.Printf("测试成功：求两个数的平均值。结果为：(%v + %v) / 2 = %v", param.Num1, param.Num2, data.Result)
			}
		})

	})

}
