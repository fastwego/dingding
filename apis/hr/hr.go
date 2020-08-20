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

// Package hr 智能人事
package hr

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiEmployeeList           = "/topapi/smartwork/hrm/employee/list"
	apiEmployeeUpdate         = "/topapi/smartwork/hrm/employee/update"
	apiEmployeeFieldGrouplist = "/topapi/smartwork/hrm/employee/field/grouplist"
	apiQueryPreentry          = "/topapi/smartwork/hrm/employee/querypreentry"
	apiQueryonjob             = "/topapi/smartwork/hrm/employee/queryonjob"
	apiQueryDimission         = "/topapi/smartwork/hrm/employee/querydimission"
	apiListDimission          = "/topapi/smartwork/hrm/employee/listdimission"
	apiAddPreentry            = "/topapi/smartwork/hrm/employee/addpreentry"
)

/*
获取员工花名册字段信息



See: https://ding-doc.dingtalk.com/doc#/serverapi2/rikq4i

POST https://oapi.dingtalk.com/topapi/smartwork/hrm/employee/list?access_token=ACCESS_TOKEN
*/
func EmployeeList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiEmployeeList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新员工花名册



See: https://ding-doc.dingtalk.com/doc#/serverapi2/aw6wuy

POST https://oapi.dingtalk.com/topapi/smartwork/hrm/employee/update?access_token=ACCESS_TOKEN
*/
func EmployeeUpdate(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiEmployeeUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询花名册元数据



See: https://ding-doc.dingtalk.com/doc#/serverapi2/beo741

POST https://oapi.dingtalk.com/topapi/smartwork/hrm/employee/field/grouplist?access_token=ACCESS_TOKEN
*/
func EmployeeFieldGrouplist(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiEmployeeFieldGrouplist, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询企业待入职员工列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/yv3mcy

POST https://oapi.dingtalk.com/topapi/smartwork/hrm/employee/querypreentry?access_token=ACCESS_TOKEN
*/
func QueryPreentry(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiQueryPreentry, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询企业在职员工列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/rafx8t

POST https://oapi.dingtalk.com/topapi/smartwork/hrm/employee/queryonjob?access_token=ACCESS_TOKEN
*/
func Queryonjob(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiQueryonjob, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询企业离职员工列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/oe0nlx

POST https://oapi.dingtalk.com/topapi/smartwork/hrm/employee/querydimission?access_token=ACCESS_TOKEN
*/
func QueryDimission(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiQueryDimission, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取离职员工离职信息



See: https://ding-doc.dingtalk.com/doc#/serverapi2/fbtugn

POST https://oapi.dingtalk.com/topapi/smartwork/hrm/employee/listdimission?access_token=ACCESS_TOKEN
*/
func ListDimission(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListDimission, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
添加企业待入职员工



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dhlu4d

POST https://oapi.dingtalk.com/topapi/smartwork/hrm/employee/addpreentry?access_token=ACCESS_TOKEN
*/
func AddPreentry(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddPreentry, bytes.NewReader(payload), "application/json;charset=utf-8")
}
