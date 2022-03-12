// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  xincan  2021/2/7 16:10
// @Update  xincan  2021/2/7 16:10
package goconvey

import (
	_ "github.com/gavv/httpexpect/v2"
	_ "github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"xincan.com.cn/istorm-cnbr-operator/kubernetes/handler"
)

// mock handler 测试
func TestMockHandlerFromKubernetesService(t *testing.T) {

	Convey("两个数处理,单元测试--handler", t, func() {
		handler.AddKubernetesHandler()
		handler.SubKubernetesHandler()
		handler.MulKubernetesHandler()
		handler.DivKubernetesHandler()
		handler.AveKubernetesHandler()
	})

}
