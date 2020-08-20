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

// Package blackboard 公告
package blackboard

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiListTop10    = "/topapi/blackboard/listtopten"
	apiCategoryList = "/topapi/blackboard/category/list"
	apiListIds      = "/topapi/blackboard/listids"
	apiCreate       = "/topapi/blackboard/create"
	apiUpdate       = "/topapi/blackboard/update"
	apiDelete       = "/topapi/blackboard/delete"
	apiGet          = "/topapi/blackboard/get"
)

/*
获取用户可查看的公告



See: https://ding-doc.dingtalk.com/doc#/serverapi2/knmd16

POST https://oapi.dingtalk.com/topapi/blackboard/listtopten?access_token=ACCESS_TOKEN
*/
func ListTop10(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListTop10, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取公告分类列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/lfaanr

POST https://oapi.dingtalk.com/topapi/blackboard/category/list?access_token=ACCESS_TOKEN
*/
func CategoryList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCategoryList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取公告id列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/cgu5xo

POST https://oapi.dingtalk.com/topapi/blackboard/listids?access_token=ACCESS_TOKEN
*/
func ListIds(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListIds, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
创建公告



See: https://ding-doc.dingtalk.com/doc#/serverapi2/owbsbu

POST https://oapi.dingtalk.com/topapi/blackboard/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新公告



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pv3zqa

POST https://oapi.dingtalk.com/topapi/blackboard/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除公告



See: https://ding-doc.dingtalk.com/doc#/serverapi2/zgmbq9

POST https://oapi.dingtalk.com/topapi/blackboard/delete?access_token=ACCESS_TOKEN
*/
func Delete(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取公告详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/xxz4hb

POST https://oapi.dingtalk.com/topapi/blackboard/get?access_token=ACCESS_TOKEN
*/
func Get(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}
