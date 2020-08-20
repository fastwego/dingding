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

// Package checkin 签到
package checkin

import (
	"bytes"
	"net/url"

	"github.com/fastwego/dingding"
)

const (
	apiRecord = "/checkin/record"
	apiGet    = "/topapi/checkin/record/get"
)

/*
获取部门用户签到记录



See: https://ding-doc.dingtalk.com/doc#/serverapi2/uyr2ah

GET https://oapi.dingtalk.com/checkin/record?access_token=ACCESS_TOKENdepartment_id=1start_time=1467707227000end_time=1467707240000offset=0size=100order=asc
*/
func Record(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiRecord + "?" + params.Encode())
}

/*
获取用户签到记录



See: https://ding-doc.dingtalk.com/doc#/serverapi2/uyr2ah

POST https://oapi.dingtalk.com/topapi/checkin/record/get?access_token=ACCESS_TOKEN
*/
func Get(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}
