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

// Package microapp 应用管理
package microapp

import (
	"bytes"
	"net/url"

	"github.com/fastwego/dingding"
)

const (
	apiList             = "/microapp/list"
	apiListByUserid     = "/microapp/list_by_userid"
	apiVisibleScopes    = "/microapp/visible_scopes"
	apiSetVisibleScopes = "/microapp/set_visible_scopes"
)

/*
获取应用列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/zc304p

POST https://oapi.dingtalk.com/microapp/list?access_token=ACCESS_TOKEN
*/
func List(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取员工可见的应用列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/zc304p

GET https://oapi.dingtalk.com/microapp/list_by_userid?access_token=ACCESS_TOKENuserid=manager
*/
func ListByUserid(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiListByUserid + "?" + params.Encode())
}

/*
获取应用的可见范围



See: https://ding-doc.dingtalk.com/doc#/serverapi2/zc304p

POST https://oapi.dingtalk.com/microapp/visible_scopes?access_token=ACCESS_TOKEN
*/
func VisibleScopes(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiVisibleScopes, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
设置应用的可见范围



See: https://ding-doc.dingtalk.com/doc#/serverapi2/zc304p

POST https://oapi.dingtalk.com/microapp/set_visible_scopes?access_token=ACCESS_TOKEN
*/
func SetVisibleScopes(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetVisibleScopes, bytes.NewReader(payload), "application/json;charset=utf-8")
}
