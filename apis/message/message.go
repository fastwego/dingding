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

// Package message 消息通知
package message

import (
	"bytes"
	"net/url"

	"github.com/fastwego/dingding"
)

const (
	apiAsyncsendV2        = "/topapi/message/corpconversation/asyncsend_v2"
	apiGetSendProgress    = "/topapi/message/corpconversation/getsendprogress"
	apiGetSendResult      = "/topapi/message/corpconversation/getsendresult"
	apiRecall             = "/topapi/message/corpconversation/recall"
	apiSend               = "/chat/send"
	apiGetReadList        = "/chat/getReadList"
	apiSendToConversation = "/message/send_to_conversation"
)

/*
查询工作通知消息的发送进度



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pgoxpy

POST https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=ACCESS_TOKEN
*/
func AsyncsendV2(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAsyncsendV2, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询工作通知消息的发送进度



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pgoxpy

POST https://oapi.dingtalk.com/topapi/message/corpconversation/getsendprogress?access_token=ACCESS_TOKEN
*/
func GetSendProgress(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetSendProgress, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询工作通知消息的发送结果



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pgoxpy

POST https://oapi.dingtalk.com/topapi/message/corpconversation/getsendresult?access_token=ACCESS_TOKEN
*/
func GetSendResult(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetSendResult, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
工作通知消息撤回



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pgoxpy

POST https://oapi.dingtalk.com/topapi/message/corpconversation/recall?access_token=ACCESS_TOKEN
*/
func Recall(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiRecall, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
发送消息到企业群



See: https://ding-doc.dingtalk.com/doc#/serverapi2/isu6nk

POST https://oapi.dingtalk.com/chat/send?access_token=ACCESS_TOKEN
*/
func Send(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSend, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询群消息已读人员列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/isu6nk

GET https://oapi.dingtalk.com/chat/getReadList?access_token=ACCESS_TOKENmessageId=MESSAGEIDcursor=CURSORsize=SIZE
*/
func GetReadList(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetReadList + "?" + params.Encode())
}

/*
发送普通消息



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pm0m06

POST https://oapi.dingtalk.com/message/send_to_conversation?access_token=ACCESS_TOKEN
*/
func SendToConversation(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSendToConversation, bytes.NewReader(payload), "application/json;charset=utf-8")
}
