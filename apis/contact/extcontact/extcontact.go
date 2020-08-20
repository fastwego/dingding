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

// Package extcontact 通讯录/外部联系人管理
package extcontact

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiListLabelGroups = "/topapi/extcontact/listlabelgroups"
	apiList            = "/topapi/extcontact/list"
	apiGet             = "/topapi/extcontact/get"
	apiCreate          = "/topapi/extcontact/create"
	apiUpdate          = "/topapi/extcontact/update"
	apiDelete          = "/topapi/extcontact/delete"
)

/*
获取外部联系人标签列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/nb93oa

POST https://oapi.dingtalk.com/topapi/extcontact/listlabelgroups?access_token=ACCESS_TOKEN
*/
func ListLabelGroups(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListLabelGroups, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取外部联系人列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/nb93oa

POST https://oapi.dingtalk.com/topapi/extcontact/list?access_token=ACCESS_TOKEN
*/
func List(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取外部联系人详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/nb93oa

POST https://oapi.dingtalk.com/topapi/extcontact/get?access_token=ACCESS_TOKEN
*/
func Get(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
添加外部联系人



See: https://ding-doc.dingtalk.com/doc#/serverapi2/nb93oa

POST https://oapi.dingtalk.com/topapi/extcontact/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新外部联系人



See: https://ding-doc.dingtalk.com/doc#/serverapi2/nb93oa

POST https://oapi.dingtalk.com/topapi/extcontact/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除外部联系人



See: https://ding-doc.dingtalk.com/doc#/serverapi2/nb93oa

POST https://oapi.dingtalk.com/topapi/extcontact/delete?access_token=ACCESS_TOKEN
*/
func Delete(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}
