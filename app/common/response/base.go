package response

import (
	"context"
	"github.com/wuqianaer/go-zero-oj/app/common/consts"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type BaseResponse[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// CodeResponse 返回code
type CodeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

const (
	SuccessCode = 200
	ErrorCode   = 500
)

const (
	SuccessMsg = "ok"
)

// JsonBaseResponseCtx 暂无支持状态码和msg的自定义
func JsonBaseResponseCtx(ctx context.Context, writer http.ResponseWriter, v any) {
	httpx.OkJsonCtx(ctx, writer, buildBaseResp(v))
}

// buildBaseResp 用于组装结构体
func buildBaseResp(v any) BaseResponse[any] {
	// 设置data的类型是any
	var resp BaseResponse[any]

	// 根据同的数据进行结构体的封装
	switch data := v.(type) {
	case error:
		resp.Code = consts.Code_Error
		resp.Msg = data.Error()
	case CodeResponse:
		resp.Code = data.Code
		resp.Msg = data.Msg
	case *CodeResponse:
		resp.Code = data.Code
		resp.Msg = data.Msg
	default:
		// 通常返回的数据
		resp.Code = consts.Code_Success
		resp.Msg = SuccessMsg
		resp.Data = v
	}
	return resp
}
