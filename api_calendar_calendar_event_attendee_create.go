package lark

import (
	"context"
)

// CreateCalendarEventAttendee 批量给日程添加参与人。
//
// - 当前身份需要有日历的 writer 或 owner 权限，并且日历的类型只能为 primary 或 shared。
// - 当前身份需要是日程的组织者，或日程设置了「参与人可邀请其它参与人」权限。
// - 新添加的日程参与人必须与日程组织者在同一个企业内。
// - 使用该接口添加会议室后，会议室会进入异步的预约流程，请求结束不代表会议室预约成功，需后续再查询预约状态。
// - 每个日程最多只能有 3000 名参与人。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/create
func (r *CalendarAPI) CreateCalendarEventAttendee(ctx context.Context, request *CreateCalendarEventAttendeeReq) (*CreateCalendarEventAttendeeResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(createCalendarEventAttendeeResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "CreateCalendarEventAttendee", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type CreateCalendarEventAttendeeReq struct {
	UserIDType       *IDType                                   `query:"user_id_type" json:"-"`      // 用户 ID 类型,**示例值**："open_id",**可选值有**：,- `open_id`：用户的 open id,- `union_id`：用户的 union id,- `user_id`：用户的 user id,**默认值**：`open_id`,**当值为 `user_id`，字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	CalendarID       string                                    `path:"calendar_id" json:"-"`        // 日历 ID,**示例值**："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	EventID          string                                    `path:"event_id" json:"-"`           // 日程 ID,**示例值**："xxxxxxxxx_0"
	Attendees        []*CreateCalendarEventAttendeeReqAttendee `json:"attendees,omitempty"`         // 新增参与人列表
	NeedNotification *bool                                     `json:"need_notification,omitempty"` // 是否给参与人发送bot通知,**示例值**：false
}

type CreateCalendarEventAttendeeReqAttendee struct {
	Type            *CalendarEventAttendeeType `json:"type,omitempty"`              // 参与人类型,**示例值**："user",**可选值有**：,- `user`：用户,- `chat`：群组,- `resource`：会议室,- `third_party`：邮箱
	IsOptional      *bool                      `json:"is_optional,omitempty"`       // 参与人是否为「可选参加」，无法编辑群参与人的此字段,**示例值**：true,**默认值**：`false`
	UserID          *string                    `json:"user_id,omitempty"`           // 参与人的用户id，依赖于user_id_type返回对应的取值，当is_external为true时，此字段只会返回open_id或者union_id,**示例值**："ou_xxxxxxxx"
	ChatID          *string                    `json:"chat_id,omitempty"`           // chat类型参与人的群组chat_id,**示例值**："om_xxxxxxxxx"
	RoomID          *string                    `json:"room_id,omitempty"`           // resource类型参与人的会议室room_id,**示例值**："omm_xxxxxxxx"
	ThirdPartyEmail *string                    `json:"third_party_email,omitempty"` // third_party类型参与人的邮箱,**示例值**："wangwu@email.com"
}

type createCalendarEventAttendeeResp struct {
	Code int                              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *CreateCalendarEventAttendeeResp `json:"data,omitempty"` // \-
}

type CreateCalendarEventAttendeeResp struct {
	Attendees []*CreateCalendarEventAttendeeRespAttendee `json:"attendees,omitempty"` // 新增参与人后的日程所有参与人列表
}

type CreateCalendarEventAttendeeRespAttendee struct {
	Type            CalendarEventAttendeeType                            `json:"type,omitempty"`              // 参与人类型,**可选值有**：,- `user`：用户,- `chat`：群组,- `resource`：会议室,- `third_party`：邮箱
	AttendeeID      string                                               `json:"attendee_id,omitempty"`       // 参与人ID
	RsvpStatus      string                                               `json:"rsvp_status,omitempty"`       // 参与人RSVP状态,**可选值有**：,- `needs_action`：参与人尚未回复状态，或表示会议室预约中,- `accept`：参与人回复接受，或表示会议室预约成功,- `tentative`：参与人回复待定,- `decline`：参与人回复拒绝，或表示会议室预约失败,- `removed`：参与人或会议室已经从日程中被移除
	IsOptional      bool                                                 `json:"is_optional,omitempty"`       // 参与人是否为「可选参加」，无法编辑群参与人的此字段,**默认值**：`false`
	IsOrganizer     bool                                                 `json:"is_organizer,omitempty"`      // 参与人是否为日程组织者
	IsExternal      bool                                                 `json:"is_external,omitempty"`       // 参与人是否为外部参与人；外部参与人不支持编辑
	DisplayName     string                                               `json:"display_name,omitempty"`      // 参与人名称
	ChatMembers     []*CreateCalendarEventAttendeeRespAttendeeChatMember `json:"chat_members,omitempty"`      // 群中的群成员，当type为Chat时有效；群成员不支持编辑
	UserID          string                                               `json:"user_id,omitempty"`           // 参与人的用户id，依赖于user_id_type返回对应的取值，当is_external为true时，此字段只会返回open_id或者union_id
	ChatID          string                                               `json:"chat_id,omitempty"`           // chat类型参与人的群组chat_id
	RoomID          string                                               `json:"room_id,omitempty"`           // resource类型参与人的会议室room_id
	ThirdPartyEmail string                                               `json:"third_party_email,omitempty"` // third_party类型参与人的邮箱
}

type CreateCalendarEventAttendeeRespAttendeeChatMember struct {
	RsvpStatus  string `json:"rsvp_status,omitempty"`  // 参与人RSVP状态,**可选值有**：,- `needs_action`：参与人尚未回复状态，或表示会议室预约中,- `accept`：参与人回复接受，或表示会议室预约成功,- `tentative`：参与人回复待定,- `decline`：参与人回复拒绝，或表示会议室预约失败,- `removed`：参与人或会议室已经从日程中被移除
	IsOptional  bool   `json:"is_optional,omitempty"`  // 参与人是否为「可选参加」,**默认值**：`false`
	DisplayName string `json:"display_name,omitempty"` // 参与人名称
	IsOrganizer bool   `json:"is_organizer,omitempty"` // 参与人是否为日程组织者
	IsExternal  bool   `json:"is_external,omitempty"`  // 参与人是否为外部参与人
}
