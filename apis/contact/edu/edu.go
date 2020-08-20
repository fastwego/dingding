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

// Package edu 家校通讯录
package edu

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiDeptList         = "/topapi/edu/dept/list"
	apiDeptGet          = "/topapi/edu/dept/get"
	apiUserList         = "/topapi/edu/user/list"
	apiUserGet          = "/topapi/edu/user/get"
	apiUserRelationList = "/topapi/edu/user/relation/list"
	apiUserRelationGet  = "/topapi/edu/user/relation/get"
)

/*
获取部门列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/zaypkk

POST https://oapi.dingtalk.com/topapi/edu/dept/list?access_token=ACCESS_TOKEN
*/
func DeptList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDeptList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取部门详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/zaypkk

POST https://oapi.dingtalk.com/topapi/edu/dept/get?access_token=ACCESS_TOKEN
*/
func DeptGet(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDeptGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
部门完整示例



See: https://ding-doc.dingtalk.com/doc#/serverapi2/zaypkk

POST https://oapi.dingtalk.com/topapi/edu/user/list?access_token=ACCESS_TOKEN
*/
func UserList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUserList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取人员详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/zaypkk

POST https://oapi.dingtalk.com/topapi/edu/user/get?access_token=ACCESS_TOKEN
*/
func UserGet(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUserGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取人员关系列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/zaypkk

POST https://oapi.dingtalk.com/topapi/edu/user/relation/list?access_token=ACCESS_TOKEN
*/
func UserRelationList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUserRelationList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取人员关系详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/zaypkk

POST https://oapi.dingtalk.com/topapi/edu/user/relation/get?access_token=ACCESS_TOKEN
*/
func UserRelationGet(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUserRelationGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}
