// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// ResendAppTicket
//
// 飞书每隔 1 小时会给应用推送一次最新的 app_ticket，应用也可以主动调用此接口，触发飞书进行及时的重新推送。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQjNz4CN2MjL0YzM
func (r *AuthAPI) ResendAppTicket(ctx context.Context, request *ResendAppTicketReq, options ...MethodOptionFunc) (*ResendAppTicketResp, *Response, error) {
	if r.cli.mock.mockAuthResendAppTicket != nil {
		r.cli.logDebug(ctx, "[lark] Auth#ResendAppTicket mock enable")
		return r.cli.mock.mockAuthResendAppTicket(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Auth#ResendAppTicket call api")
	r.cli.logDebug(ctx, "[lark] Auth#ResendAppTicket request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:       "POST",
		URL:          "https://open.feishu.cn/open-apis/auth/v3/app_ticket/resend/",
		Body:         request,
		MethodOption: newMethodOption(options),
	}
	resp := new(resendAppTicketResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Auth#ResendAppTicket POST https://open.feishu.cn/open-apis/auth/v3/app_ticket/resend/ failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Auth#ResendAppTicket POST https://open.feishu.cn/open-apis/auth/v3/app_ticket/resend/ failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Auth", "ResendAppTicket", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Auth#ResendAppTicket request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockAuthResendAppTicket(f func(ctx context.Context, request *ResendAppTicketReq, options ...MethodOptionFunc) (*ResendAppTicketResp, *Response, error)) {
	r.mockAuthResendAppTicket = f
}

func (r *Mock) UnMockAuthResendAppTicket() {
	r.mockAuthResendAppTicket = nil
}

type ResendAppTicketReq struct {
	AppID     string `json:"app_id,omitempty"`     // 应用唯一标识，创建应用后获得
	AppSecret string `json:"app_secret,omitempty"` // 应用秘钥，创建应用后获得
}

type resendAppTicketResp struct {
	Code int                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *ResendAppTicketResp `json:"data,omitempty"`
}

type ResendAppTicketResp struct{}