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

// Package calendar 日程
package calendar

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiCreate         = "/topapi/calendar/v2/event/create"
	apiUpdate         = "/topapi/calendar/v2/event/update"
	apiCancel         = "/topapi/calendar/v2/event/cancel"
	apiAttendeeUpdate = "/topapi/calendar/v2/attendee/update"
)

/*
创建日程



See: https://ding-doc.dingtalk.com/doc#/serverapi2/iqel76

POST https://oapi.dingtalk.com/topapi/calendar/v2/event/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改日程



See: https://ding-doc.dingtalk.com/doc#/serverapi2/zg9qoo

POST https://oapi.dingtalk.com/topapi/calendar/v2/event/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
取消日程



See: https://ding-doc.dingtalk.com/doc#/serverapi2/nf4c9z

POST https://oapi.dingtalk.com/topapi/calendar/v2/event/cancel?access_token=ACCESS_TOKEN
*/
func Cancel(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCancel, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
新增/删除参会人



See: https://ding-doc.dingtalk.com/doc#/serverapi2/atuhpy

POST https://oapi.dingtalk.com/topapi/calendar/v2/attendee/update?access_token=ACCESS_TOKEN
*/
func AttendeeUpdate(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAttendeeUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}
