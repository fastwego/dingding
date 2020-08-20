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

// Package callback 业务事件回调
package callback

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiRegisterCallBack        = "/call_back/register_call_back"
	apiGetCallBack             = "/call_back/get_call_back"
	apiUpdateCallBack          = "/call_back/update_call_back"
	apiDeleteCallBack          = "/call_back/delete_call_back"
	apiGetCallBackFailedResult = "/call_back/get_call_back_failed_result"
)

/*
注册业务事件回调接口



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pwz3r5

POST https://oapi.dingtalk.com/call_back/register_call_back?access_token=ACCESS_TOKEN
*/
func RegisterCallBack(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiRegisterCallBack, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
测试回调URL



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pwz3r5

GET https://oapi.dingtalk.com/call_back/get_call_back?access_token=ACCESS_TOKEN
*/
func GetCallBack(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetCallBack)
}

/*
更新事件回调接口



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pwz3r5

POST https://oapi.dingtalk.com/call_back/update_call_back?access_token=ACCESS_TOKEN
*/
func UpdateCallBack(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateCallBack, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除事件回调接口



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pwz3r5

GET https://oapi.dingtalk.com/call_back/delete_call_back?access_token=ACCESS_TOKEN
*/
func DeleteCallBack(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiDeleteCallBack)
}

/*
获取回调失败的结果



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pwz3r5

GET https://oapi.dingtalk.com/call_back/get_call_back_failed_result?access_token=ACCESS_TOKEN
*/
func GetCallBackFailedResult(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetCallBackFailedResult)
}
