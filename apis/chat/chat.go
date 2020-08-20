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

// Package chat 群会话管理
package chat

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiCreate             = "/chat/create"
	apiUpdate             = "/chat/update"
	apiGet                = "/chat/get"
	apiSubadminUpdate     = "/topapi/chat/subadmin/update"
	apiFriendSwitchUpdate = "/topapi/chat/member/friendswitch/update"
)

/*
创建会话



See: https://ding-doc.dingtalk.com/doc#/serverapi2/fg9dze

POST https://oapi.dingtalk.com/chat/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改会话



See: https://ding-doc.dingtalk.com/doc#/serverapi2/lewq17

POST https://oapi.dingtalk.com/chat/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取会话



See: https://ding-doc.dingtalk.com/doc#/serverapi2/kruhy0

GET https://oapi.dingtalk.com/chat/get?access_token=ACCESS_TOKEN
*/
func Get(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGet)
}

/*
设置群管理员



See: https://ding-doc.dingtalk.com/doc#/serverapi2/leqbe8

POST https://oapi.dingtalk.com/topapi/chat/subadmin/update?access_token=ACCESS_TOKEN
*/
func SubadminUpdate(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSubadminUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
设置禁止群成员私聊



See: https://ding-doc.dingtalk.com/doc#/serverapi2/zb74qt

POST https://oapi.dingtalk.com/topapi/chat/member/friendswitch/update?access_token=ACCESS_TOKEN
*/
func FriendSwitchUpdate(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiFriendSwitchUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}
