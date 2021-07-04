// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetDriveMemberPermissionList 该接口用于根据 filetoken 查询协作者，目前包括人("user")和群("chat") 。
//
// 你能获取到协作者列表的前提是你对该文档有分享权限
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATN3UjLwUzN14CM1cTN
func (r *DriveService) GetDriveMemberPermissionList(ctx context.Context, request *GetDriveMemberPermissionListReq, options ...MethodOptionFunc) (*GetDriveMemberPermissionListResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveMemberPermissionList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDriveMemberPermissionList mock enable")
		return r.cli.mock.mockDriveGetDriveMemberPermissionList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDriveMemberPermissionList",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/drive/permission/member/list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDriveMemberPermissionListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveGetDriveMemberPermissionList(f func(ctx context.Context, request *GetDriveMemberPermissionListReq, options ...MethodOptionFunc) (*GetDriveMemberPermissionListResp, *Response, error)) {
	r.mockDriveGetDriveMemberPermissionList = f
}

func (r *Mock) UnMockDriveGetDriveMemberPermissionList() {
	r.mockDriveGetDriveMemberPermissionList = nil
}

type GetDriveMemberPermissionListReq struct {
	Token string `json:"token,omitempty"` // 文件的 token，获取方式见 [对接前说明](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 4 项
	Type  string `json:"type,omitempty"`  // 文档类型  "doc"  or  "sheet" or "file"
}

type getDriveMemberPermissionListResp struct {
	Code int64                             `json:"code,omitempty"`
	Msg  string                            `json:"msg,omitempty"`
	Data *GetDriveMemberPermissionListResp `json:"data,omitempty"`
}

type GetDriveMemberPermissionListResp struct {
	Members []*GetDriveMemberPermissionListRespMember `json:"members,omitempty"` // 协作者列表
}

type GetDriveMemberPermissionListRespMember struct {
	MemberType   string `json:"member_type,omitempty"`    // 协作者类型 "user" or "chat"
	MemberOpenID string `json:"member_open_id,omitempty"` // 协作者openid
	MemberUserID string `json:"member_user_id,omitempty"` // 协作者userid(仅当member_type="user"时有效)
	Perm         string `json:"perm,omitempty"`           // 协作者权限 (注意: **有"edit"权限的协作者一定有"view"权限**)
}
