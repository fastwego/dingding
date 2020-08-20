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

// Package attendance 考勤
package attendance

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiListSchedule                   = "/topapi/attendance/listschedule"
	apiListByDay                      = "/topapi/attendance/schedule/listbyday"
	apiListByUsers                    = "/topapi/attendance/schedule/listbyusers"
	apiScheduleAsync                  = "/topapi/attendance/group/schedule/async"
	apiListByIds                      = "/topapi/attendance/schedule/result/listbyids"
	apiAttendanceShiftList            = "/topapi/attendance/shift/list"
	apiAttendanceShiftQuery           = "/topapi/attendance/shift/query"
	apiAttendanceShiftSearch          = "/topapi/attendance/shift/search"
	apiListRecord                     = "/attendance/listRecord"
	apiAttendanceList                 = "/attendance/list"
	apiGetLeaveStatus                 = "/topapi/attendance/getleavestatus"
	apiGetSimpleGroups                = "/topapi/attendance/getsimplegroups"
	apiAttendanceGroupMinimalismList  = "/topapi/attendance/group/minimalism/list"
	apiAttendanceGroupSearch          = "/topapi/attendance/group/search"
	apiAttendanceGroupQuery           = "/topapi/attendance/group/query"
	apiGetUserGroup                   = "/topapi/attendance/getusergroup"
	apiAttendanceGroupMemberUsersList = "/topapi/attendance/group/memberusers/list"
	apiAttendanceGroupMemberListbyids = "/topapi/attendance/group/member/listbyids"
	apiAttendanceGroupMemberUpdate    = "/topapi/attendance/group/member/update"
	apiAttendanceGroupMemberList      = "/topapi/attendance/group/member/list"
	apiIsOpenSmartReport              = "/topapi/attendance/isopensmartreport"
	apiGetAttColumns                  = "/topapi/attendance/getattcolumns"
	apiGetColumnVal                   = "/topapi/attendance/getcolumnval"
	apiGetLeaveTimeByNames            = "/topapi/attendance/getleavetimebynames"
)

/*
查询企业考勤排班详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ufc8dl

POST https://oapi.dingtalk.com/topapi/attendance/listschedule?access_token=ACCESS_TOKEN
*/
func ListSchedule(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListSchedule, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询成员排班信息



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ufc8dl

POST https://oapi.dingtalk.com/topapi/attendance/schedule/listbyday?access_token=ACCESS_TOKEN
*/
func ListByDay(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListByDay, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量查询成员排班信息



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ufc8dl

POST https://oapi.dingtalk.com/topapi/attendance/schedule/listbyusers?access_token=ACCESS_TOKEN
*/
func ListByUsers(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListByUsers, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
排班制考勤组排班



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ufc8dl

POST https://oapi.dingtalk.com/topapi/attendance/group/schedule/async?access_token=ACCESS_TOKEN
*/
func ScheduleAsync(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiScheduleAsync, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询排班打卡结果



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ufc8dl

POST https://oapi.dingtalk.com/topapi/attendance/schedule/result/listbyids?access_token=ACCESS_TOKEN
*/
func ListByIds(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListByIds, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量查询班次摘要信息



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pmug9c

POST https://oapi.dingtalk.com/topapi/attendance/shift/list?access_token=ACCESS_TOKEN
*/
func AttendanceShiftList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAttendanceShiftList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询班次详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pmug9c

POST https://oapi.dingtalk.com/topapi/attendance/shift/query?access_token=ACCESS_TOKEN
*/
func AttendanceShiftQuery(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAttendanceShiftQuery, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
按名称搜索班次



See: https://ding-doc.dingtalk.com/doc#/serverapi2/pmug9c

POST https://oapi.dingtalk.com/topapi/attendance/shift/search?access_token=ACCESS_TOKEN
*/
func AttendanceShiftSearch(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAttendanceShiftSearch, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取打卡详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/potcn9

POST https://oapi.dingtalk.com/attendance/listRecord?access_token=ACCESS_TOKEN
*/
func ListRecord(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListRecord, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取打卡结果



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ul33mm

POST https://oapi.dingtalk.com/attendance/list?access_token=ACCESS_TOKEN
*/
func AttendanceList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAttendanceList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询请假状态



See: https://ding-doc.dingtalk.com/doc#/serverapi2/chhwzv

POST https://oapi.dingtalk.com/topapi/attendance/getleavestatus?access_token=ACCESS_TOKEN
*/
func GetLeaveStatus(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetLeaveStatus, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量获取企业考勤组详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ksk4o2

POST https://oapi.dingtalk.com/topapi/attendance/getsimplegroups?access_token=ACCESS_TOKEN
*/
func GetSimpleGroups(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetSimpleGroups, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取考勤组摘要



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ksk4o2

POST https://oapi.dingtalk.com/topapi/attendance/group/minimalism/list?access_token=ACCESS_TOKEN
*/
func AttendanceGroupMinimalismList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAttendanceGroupMinimalismList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
搜索考勤组摘要



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ksk4o2

POST https://oapi.dingtalk.com/topapi/attendance/group/search?access_token=ACCESS_TOKEN
*/
func AttendanceGroupSearch(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAttendanceGroupSearch, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取考勤组详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ksk4o2

POST https://oapi.dingtalk.com/topapi/attendance/group/query?access_token=ACCESS_TOKEN
*/
func AttendanceGroupQuery(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAttendanceGroupQuery, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户考勤组



See: https://ding-doc.dingtalk.com/doc#/serverapi2/uvt47u

POST https://oapi.dingtalk.com/topapi/attendance/getusergroup?access_token=ACCESS_TOKEN
*/
func GetUserGroup(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUserGroup, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取考勤组员工id



See: https://ding-doc.dingtalk.com/doc#/serverapi2/uvt47u

POST https://oapi.dingtalk.com/topapi/attendance/group/memberusers/list?access_token=ACCESS_TOKEN
*/
func AttendanceGroupMemberUsersList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAttendanceGroupMemberUsersList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
考勤组成员校验



See: https://ding-doc.dingtalk.com/doc#/serverapi2/uvt47u

POST https://oapi.dingtalk.com/topapi/attendance/group/member/listbyids?access_token=ACCESS_TOKEN
*/
func AttendanceGroupMemberListbyids(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAttendanceGroupMemberListbyids, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
考勤组成员更新



See: https://ding-doc.dingtalk.com/doc#/serverapi2/uvt47u

POST https://oapi.dingtalk.com/topapi/attendance/group/member/update?access_token=ACCESS_TOKEN
*/
func AttendanceGroupMemberUpdate(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAttendanceGroupMemberUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取考勤组成员



See: https://ding-doc.dingtalk.com/doc#/serverapi2/uvt47u

POST https://oapi.dingtalk.com/topapi/attendance/group/member/list?access_token=ACCESS_TOKEN
*/
func AttendanceGroupMemberList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAttendanceGroupMemberList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
是否启用智能统计报表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vg7ned

POST https://oapi.dingtalk.com/topapi/attendance/isopensmartreport?access_token=ACCESS_TOKEN
*/
func IsOpenSmartReport(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiIsOpenSmartReport, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取报表列定义



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vg7ned

POST https://oapi.dingtalk.com/topapi/attendance/getattcolumns?access_token=ACCESS_TOKEN
*/
func GetAttColumns(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetAttColumns, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取报表列值



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vg7ned

POST https://oapi.dingtalk.com/topapi/attendance/getcolumnval?access_token=ACCESS_TOKEN
*/
func GetColumnVal(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetColumnVal, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取报表假期数据



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vg7ned

POST https://oapi.dingtalk.com/topapi/attendance/getleavetimebynames?access_token=ACCESS_TOKEN
*/
func GetLeaveTimeByNames(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetLeaveTimeByNames, bytes.NewReader(payload), "application/json;charset=utf-8")
}
