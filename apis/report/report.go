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

// Package report 日志
package report

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiList               = "/topapi/report/list"
	apiSimplelist         = "/topapi/report/simplelist"
	apiStatistics         = "/topapi/report/statistics"
	apiListByType         = "/topapi/report/statistics/listbytype"
	apiReportReceiverList = "/topapi/report/receiver/list"
	apiReportCommentList  = "/topapi/report/comment/list"
	apiGetUnreadCount     = "/topapi/report/getunreadcount"
	apiListByUserId       = "/topapi/report/template/listbyuserid"
)

/*
获取企业用户日志数据



See: https://ding-doc.dingtalk.com/doc#/serverapi2/yknhmg

POST https://oapi.dingtalk.com/topapi/report/list?access_token=ACCESS_TOKEN
*/
func List(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户发送日志概要信息



See: https://ding-doc.dingtalk.com/doc#/serverapi2/cp4mxh

POST https://oapi.dingtalk.com/topapi/report/simplelist?access_token=ACCESS_TOKEN
*/
func Simplelist(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSimplelist, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取日志统计数据



See: https://ding-doc.dingtalk.com/doc#/serverapi2/lhlow3

POST https://oapi.dingtalk.com/topapi/report/statistics?access_token=ACCESS_TOKEN
*/
func Statistics(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiStatistics, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取日志相关人员列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/qga9wa

POST https://oapi.dingtalk.com/topapi/report/statistics/listbytype?access_token=ACCESS_TOKEN
*/
func ListByType(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListByType, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取日志接收人员列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/el6lgx

POST https://oapi.dingtalk.com/topapi/report/receiver/list?access_token=ACCESS_TOKEN
*/
func ReportReceiverList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiReportReceiverList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取日志评论详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pmyeg8

POST https://oapi.dingtalk.com/topapi/report/comment/list?access_token=ACCESS_TOKEN
*/
func ReportCommentList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiReportCommentList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户日志未读数



See: https://ding-doc.dingtalk.com/doc#/serverapi2/yz2vbx

POST https://oapi.dingtalk.com/topapi/report/getunreadcount?access_token=ACCESS_TOKEN
*/
func GetUnreadCount(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUnreadCount, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户可见的日志模板



See: https://ding-doc.dingtalk.com/doc#/serverapi2/lbt9b4

POST https://oapi.dingtalk.com/topapi/report/template/listbyuserid?access_token=ACCESS_TOKEN
*/
func ListByUserId(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListByUserId, bytes.NewReader(payload), "application/json;charset=utf-8")
}
