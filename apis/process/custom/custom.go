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

// Package custom 智能工作流/自有
package custom

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiSave            = "/topapi/process/save"
	apiGetByName       = "/topapi/process/get_by_name"
	apiDelete          = "/topapi/process/delete"
	apiCreate          = "/topapi/process/workrecord/create"
	apiUpdate          = "/topapi/process/workrecord/update"
	apiBatchUpdate     = "/topapi/process/workrecord/batchupdate"
	apiTaskCreate      = "/topapi/process/workrecord/task/create"
	apiTaskUpdate      = "/topapi/process/workrecord/task/update"
	apiTaskGroupCancel = "/topapi/process/workrecord/taskgroup/cancel"
	apiTaskQuery       = "/topapi/process/workrecord/task/query"
)

/*
创建或更新模板



See: https://ding-doc.dingtalk.com/doc#/serverapi2/xcgg03

POST https://oapi.dingtalk.com/topapi/process/save?access_token=ACCESS_TOKEN
*/
func Save(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSave, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取模板code



See: https://ding-doc.dingtalk.com/doc#/serverapi2/uts09n

POST https://oapi.dingtalk.com/topapi/process/get_by_name?access_token=ACCESS_TOKEN
*/
func GetByName(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetByName, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除模板



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ilni1r

POST https://oapi.dingtalk.com/topapi/process/delete?access_token=ACCESS_TOKEN
*/
func Delete(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
创建实例



See: https://ding-doc.dingtalk.com/doc#/serverapi2/xfg2g9

POST https://oapi.dingtalk.com/topapi/process/workrecord/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新实例状态



See: https://ding-doc.dingtalk.com/doc#/serverapi2/as4xlh

POST https://oapi.dingtalk.com/topapi/process/workrecord/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量更新实例状态



See: https://ding-doc.dingtalk.com/doc#/serverapi2/as4xlh

POST https://oapi.dingtalk.com/topapi/process/workrecord/batchupdate?access_token=ACCESS_TOKEN
*/
func BatchUpdate(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
创建待办事项



See: https://ding-doc.dingtalk.com/doc#/serverapi2/rbe7gs

POST https://oapi.dingtalk.com/topapi/process/workrecord/task/create?access_token=ACCESS_TOKEN
*/
func TaskCreate(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTaskCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新待办状态



See: https://ding-doc.dingtalk.com/doc#/serverapi2/riycru

POST https://oapi.dingtalk.com/topapi/process/workrecord/task/update?access_token=ACCESS_TOKEN
*/
func TaskUpdate(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTaskUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量取消待办



See: https://ding-doc.dingtalk.com/doc#/serverapi2/yaiekm

POST https://oapi.dingtalk.com/topapi/process/workrecord/taskgroup/cancel?access_token=ACCESS_TOKEN
*/
func TaskGroupCancel(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTaskGroupCancel, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询用户待办列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/riueql

POST https://oapi.dingtalk.com/topapi/process/workrecord/task/query?access_token=ACCESS_TOKEN
*/
func TaskQuery(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTaskQuery, bytes.NewReader(payload), "application/json;charset=utf-8")
}
