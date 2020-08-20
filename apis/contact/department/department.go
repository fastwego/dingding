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

// Package department 通讯录/部门管理
package department

import (
	"bytes"
	"net/url"

	"github.com/fastwego/dingding"
)

const (
	apiCreate                = "/department/create"
	apiUpdate                = "/department/update"
	apiDelete                = "/department/delete"
	apiListIds               = "/department/list_ids"
	apiList                  = "/department/list"
	apiGet                   = "/department/get"
	apiListParentDeptsByDept = "/department/list_parent_depts_by_dept"
	apiListParentDepts       = "/department/list_parent_depts"
)

/*
创建部门



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dubakq

POST https://oapi.dingtalk.com/department/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新部门



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dubakq

POST https://oapi.dingtalk.com/department/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除部门



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dubakq

GET https://oapi.dingtalk.com/department/delete?access_token=ACCESS_TOKENid=ID
*/
func Delete(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiDelete + "?" + params.Encode())
}

/*
获取子部门ID列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dubakq

GET https://oapi.dingtalk.com/department/list_ids?access_token=ACCESS_TOKENid=ID
*/
func ListIds(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiListIds + "?" + params.Encode())
}

/*
获取部门列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dubakq

GET https://oapi.dingtalk.com/department/list?access_token=ACCESS_TOKEN
*/
func List(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiList)
}

/*
获取部门详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dubakq

GET https://oapi.dingtalk.com/department/get?access_token=ACCESS_TOKENid=123
*/
func Get(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGet + "?" + params.Encode())
}

/*
查询部门的所有上级父部门路径



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dubakq

GET https://oapi.dingtalk.com/department/list_parent_depts_by_dept?access_token=ACCESS_TOKENid=ID
*/
func ListParentDeptsByDept(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiListParentDeptsByDept + "?" + params.Encode())
}

/*
查询指定用户的所有上级父部门路径



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dubakq

GET https://oapi.dingtalk.com/department/list_parent_depts?access_token=ACCESS_TOKENuserId=USERID
*/
func ListParentDepts(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiListParentDepts + "?" + params.Encode())
}
