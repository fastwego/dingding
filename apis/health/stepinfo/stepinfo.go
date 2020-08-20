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

// Package stepinfo 钉钉运动
package stepinfo

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiGetUserStatus = "/topapi/health/stepinfo/getuserstatus"
	apiList          = "/topapi/health/stepinfo/list"
	apiListByUserid  = "/topapi/health/stepinfo/listbyuserid"
)

/*
获取用户钉钉运动开启状态



See: https://ding-doc.dingtalk.com/doc#/serverapi2/emdh5f

POST https://oapi.dingtalk.com/topapi/health/stepinfo/getuserstatus?access_token=ACCESS_TOKEN
*/
func GetUserStatus(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUserStatus, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取个人或部门的钉钉运动数据



See: https://ding-doc.dingtalk.com/doc#/serverapi2/emdh5f

POST https://oapi.dingtalk.com/topapi/health/stepinfo/list?access_token=ACCESS_TOKEN
*/
func List(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量获取钉钉运动数据



See: https://ding-doc.dingtalk.com/doc#/serverapi2/emdh5f

POST https://oapi.dingtalk.com/topapi/health/stepinfo/listbyuserid?access_token=ACCESS_TOKEN
*/
func ListByUserid(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListByUserid, bytes.NewReader(payload), "application/json;charset=utf-8")
}
