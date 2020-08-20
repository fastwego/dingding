// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package datav 专属钉钉/数据统计
package datav

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiGetAccessToken              = "/topapi/kac/datav/videolive/get_access_token=ACCESS_TOKEN"
	apiKacDatavDeptVideoliveList   = "/topapi/kac/datav/dept/videolive/list"
	apiKacDatavVideoliveDetailList = "/topapi/kac/datav/videolive/detail/list"
	apiKacDatavVideoliveViewerList = "/topapi/kac/datav/videolive/viewer/list"
	apiKacDatavGroupGet            = "/topapi/kac/datav/group/get"
	apiKacV2DatavVideoconfGet      = "/topapi/kac/v2/datav/videoconf/get"
	apiKacDatavDeptVideoconfList   = "/topapi/kac/datav/dept/videoconf/list"
	apiKacDatavVideoconfDetailList = "/topapi/kac/datav/videoconf/detail/list"
	apiKacDatavTelconfGet          = "/topapi/kac/datav/telconf/get"
	apiKacDatavDeptTelconfList     = "/topapi/kac/datav/dept/telconf/list"
	apiKacDatavTelconfDetailList   = "/topapi/kac/datav/telconf/detail/list"
	apiKacDatavChatSummaryGet      = "/topapi/kac/datav/chat/summary/get"
	apiKacDatavMicroappDetailList  = "/topapi/kac/datav/microapp/detail/list"
	apiKacDatavDauSummaryGet       = "/topapi/kac/datav/dau/summary/get"
	apiKacDatavDeptDauList         = "/topapi/kac/datav/dept/dau/list"
	apiKacDatavInactivatedUserList = "/topapi/kac/datav/inactivated/user/list"
)

/*
使用说明



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/videolive/get_access_token=ACCESS_TOKEN
*/
func GetAccessToken(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetAccessToken, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业视频直播统计列表（部门维度）



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/dept/videolive/list?access_token=ACCESS_TOKEN
*/
func KacDatavDeptVideoliveList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacDatavDeptVideoliveList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业视频直播情况



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/videolive/detail/list?access_token=ACCESS_TOKEN
*/
func KacDatavVideoliveDetailList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacDatavVideoliveDetailList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业视频直播观看人员明细列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/videolive/viewer/list?access_token=ACCESS_TOKEN
*/
func KacDatavVideoliveViewerList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacDatavVideoliveViewerList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业各类群组创建情况



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/group/get?access_token=ACCESS_TOKEN
*/
func KacDatavGroupGet(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacDatavGroupGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业某天的视频会议统计数据2.0



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/v2/datav/videoconf/get?access_token=ACCESS_TOKEN
*/
func KacV2DatavVideoconfGet(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacV2DatavVideoconfGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业某天的所有部门视频会议统计数据



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/dept/videoconf/list?access_token=ACCESS_TOKEN
*/
func KacDatavDeptVideoconfList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacDatavDeptVideoconfList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业视频会议明细列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/videoconf/detail/list?access_token=ACCESS_TOKEN
*/
func KacDatavVideoconfDetailList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacDatavVideoconfDetailList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业某天的电话会议数据



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/telconf/get?access_token=ACCESS_TOKEN
*/
func KacDatavTelconfGet(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacDatavTelconfGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业某天的所有部门电话会议统计列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/dept/telconf/list?access_token=ACCESS_TOKEN
*/
func KacDatavDeptTelconfList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacDatavDeptTelconfList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业电话会议明细列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/telconf/detail/list?access_token=ACCESS_TOKEN
*/
func KacDatavTelconfDetailList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacDatavTelconfDetailList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业聊天数据



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/chat/summary/get?access_token=ACCESS_TOKEN
*/
func KacDatavChatSummaryGet(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacDatavChatSummaryGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业应用访问情况



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/microapp/detail/list?access_token=ACCESS_TOKEN
*/
func KacDatavMicroappDetailList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacDatavMicroappDetailList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
统计企业活跃用户



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/dau/summary/get?access_token=ACCESS_TOKEN
*/
func KacDatavDauSummaryGet(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacDatavDauSummaryGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
企业活跃用户统计列表（部门维度）



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/dept/dau/list?access_token=ACCESS_TOKEN
*/
func KacDatavDeptDauList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacDatavDeptDauList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
企业通讯录未激活用户列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vbywnn

POST https://oapi.dingtalk.com/topapi/kac/datav/inactivated/user/list?access_token=ACCESS_TOKEN
*/
func KacDatavInactivatedUserList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiKacDatavInactivatedUserList, bytes.NewReader(payload), "application/json;charset=utf-8")
}
