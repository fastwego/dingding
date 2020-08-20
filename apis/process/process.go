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

// Package process 智能工作流/官方
package process

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiCreate        = "/topapi/processinstance/create"
	apiListIds       = "/topapi/processinstance/listids"
	apiGet           = "/topapi/processinstance/get"
	apiGetTodoNum    = "/topapi/process/gettodonum"
	apiListByUserId  = "/topapi/process/listbyuserid"
	apiCspaceInfo    = "/topapi/processinstance/cspace/info"
	apiCspacePreview = "/topapi/processinstance/cspace/preview"
)

/*
发起审批实例



See: https://ding-doc.dingtalk.com/doc#/serverapi2/cmct1a

POST https://oapi.dingtalk.com/topapi/processinstance/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量获取审批实例id



See: https://ding-doc.dingtalk.com/doc#/serverapi2/hh8lx5

POST https://oapi.dingtalk.com/topapi/processinstance/listids?access_token=ACCESS_TOKEN
*/
func ListIds(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListIds, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取审批实例详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/xgqkvx

POST https://oapi.dingtalk.com/topapi/processinstance/get?access_token=ACCESS_TOKEN
*/
func Get(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户待审批数量



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ui5305

POST https://oapi.dingtalk.com/topapi/process/gettodonum?access_token=ACCESS_TOKEN
*/
func GetTodoNum(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetTodoNum, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户可见的审批模板



See: https://ding-doc.dingtalk.com/doc#/serverapi2/tcwmha

POST https://oapi.dingtalk.com/topapi/process/listbyuserid?access_token=ACCESS_TOKEN
*/
func ListByUserId(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListByUserId, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取审批钉盘空间信息



See: https://ding-doc.dingtalk.com/doc#/serverapi2/xq6ml3

POST https://oapi.dingtalk.com/topapi/processinstance/cspace/info?access_token=ACCESS_TOKEN
*/
func CspaceInfo(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCspaceInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
预览审批附件



See: https://ding-doc.dingtalk.com/doc#/serverapi2/sg687u

POST https://oapi.dingtalk.com/topapi/processinstance/cspace/preview?access_token=ACCESS_TOKEN
*/
func CspacePreview(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCspacePreview, bytes.NewReader(payload), "application/json;charset=utf-8")
}
