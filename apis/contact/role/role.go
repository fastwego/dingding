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

// Package role 通讯录/角色管理
package role

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiList               = "/topapi/role/list"
	apiSimpleList         = "/topapi/role/simplelist"
	apiGetRoleGroup       = "/topapi/role/getrolegroup"
	apiGetRole            = "/topapi/role/getrole"
	apiAddRole            = "/role/add_role"
	apiUpdateRole         = "/role/update_role"
	apiDeleteRole         = "/topapi/role/deleterole"
	apiAddRoleGroup       = "/role/add_role_group"
	apiAddRolesForEmps    = "/topapi/role/addrolesforemps"
	apiRemoveRolesForEmps = "/topapi/role/removerolesforemps"
	apiScopeUpdate        = "/topapi/role/scope/update"
)

/*
获取角色列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dnu5l1

POST https://oapi.dingtalk.com/topapi/role/list?access_token=ACCESS_TOKEN
*/
func List(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取角色下的员工列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dnu5l1

POST https://oapi.dingtalk.com/topapi/role/simplelist?access_token=ACCESS_TOKEN
*/
func SimpleList(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSimpleList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取角色组



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dnu5l1

POST https://oapi.dingtalk.com/topapi/role/getrolegroup?access_token=ACCESS_TOKEN
*/
func GetRoleGroup(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetRoleGroup, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取角色详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dnu5l1

POST https://oapi.dingtalk.com/topapi/role/getrole?access_token=ACCESS_TOKEN
*/
func GetRole(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetRole, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
创建角色



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dnu5l1

POST https://oapi.dingtalk.com/role/add_role?access_token=ACCESS_TOKEN
*/
func AddRole(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddRole, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新角色



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dnu5l1

POST https://oapi.dingtalk.com/role/update_role?access_token=ACCESS_TOKEN
*/
func UpdateRole(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateRole, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除角色



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dnu5l1

POST https://oapi.dingtalk.com/topapi/role/deleterole?access_token=ACCESS_TOKEN
*/
func DeleteRole(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDeleteRole, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
创建角色组



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dnu5l1

POST https://oapi.dingtalk.com/role/add_role_group?access_token=ACCESS_TOKEN
*/
func AddRoleGroup(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddRoleGroup, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量增加员工角色



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dnu5l1

POST https://oapi.dingtalk.com/topapi/role/addrolesforemps?access_token=ACCESS_TOKEN
*/
func AddRolesForEmps(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddRolesForEmps, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量删除员工角色



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dnu5l1

POST https://oapi.dingtalk.com/topapi/role/removerolesforemps?access_token=ACCESS_TOKEN
*/
func RemoveRolesForEmps(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiRemoveRolesForEmps, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
设定角色成员管理范围



See: https://ding-doc.dingtalk.com/doc#/serverapi2/dnu5l1

POST https://oapi.dingtalk.com/topapi/role/scope/update?access_token=ACCESS_TOKEN
*/
func ScopeUpdate(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiScopeUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}
