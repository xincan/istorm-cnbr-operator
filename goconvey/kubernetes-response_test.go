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
	"hatech.com.cn/istorm-cnbr-operator/utils"
	"hatech.com.cn/istorm-cnbr-operator/utils/response"
	"hatech.com.cn/istorm-cnbr-operator/vo"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var ktv = &vo.KubernetesVo{
	Result: 2,
}

type UrlInfo struct {
	Res  http.ResponseWriter
	Req  *http.Request
	Vars map[string]string
}

// mock handler 测试
func TestMockResponseFromKubernetesService(t *testing.T) {

	Convey("统一返回值测试--response", t, func() {

		Convey("统一返回解码测试", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			recorder := httptest.NewRecorder()
			err := utils.EncodeKubernetesResponse(ctx, recorder, nil)
			if err == nil {
				log.Printf("测试成功：统一返回解码测试。")
			}
		})

		Convey("统一返回成功信息日志记录函数(主要是分页)", func() {
			res := response.ResultSuccessManyAndCount(1, ktv)
			data := (res.Data).(*vo.KubernetesVo)
			if data.Result == ktv.Result {
				log.Printf(" 测试成功：统一返回成功信息日志记录函数(主要是分页)。data.Result == ktv.Result == %v", data.Result)
			}
		})

	})

}
