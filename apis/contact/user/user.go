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

// Package user 通讯录/用户管理
package user

import (
	"bytes"
	"net/url"

	"github.com/fastwego/dingding"
)

const (
	apiCreate             = "/user/create"
	apiUpdate             = "/user/update"
	apiDelete             = "/user/delete"
	apiGet                = "/user/get"
	apiGetDeptMember      = "/user/getDeptMember"
	apiSimpleList         = "/user/simplelist"
	apiListByPage         = "/user/listbypage"
	apiGetAdmin           = "/user/get_admin"
	apiGetAdminScope      = "/topapi/user/get_admin_scope"
	apiGetUseridByUnionid = "/user/getUseridByUnionid"
	apiGetByMobile        = "/user/get_by_mobile"
	apiGetOrgUserCount    = "/user/get_org_user_count"
	apiInactiveUserGet    = "/topapi/inactive/user/get"
)

/*
创建用户



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ege851

POST https://oapi.dingtalk.com/user/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新用户



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ege851

POST https://oapi.dingtalk.com/user/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除用户



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ege851

GET https://oapi.dingtalk.com/user/delete?access_token=ACCESS_TOKENuserid=zhangsan
*/
func Delete(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiDelete + "?" + params.Encode())
}

/*
获取用户详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ege851

GET https://oapi.dingtalk.com/user/get?access_token=ACCESS_TOKENuserid=zhangsan
*/
func Get(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGet + "?" + params.Encode())
}

/*
获取部门用户userid列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ege851

GET https://oapi.dingtalk.com/user/getDeptMember?access_token=ACCESS_TOKENdeptId=1
*/
func GetDeptMember(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetDeptMember + "?" + params.Encode())
}

/*
获取部门用户



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ege851

GET https://oapi.dingtalk.com/user/simplelist?access_token=ACCESS_TOKENdepartment_id=1
*/
func SimpleList(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiSimpleList + "?" + params.Encode())
}

/*
获取部门用户详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ege851

GET https://oapi.dingtalk.com/user/listbypage?access_token=ACCESS_TOKENdepartment_id=1
*/
func ListByPage(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiListByPage + "?" + params.Encode())
}

/*
获取管理员列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ege851

GET https://oapi.dingtalk.com/user/get_admin?access_token=ACCESS_TOKEN
*/
func GetAdmin(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetAdmin)
}

/*
获取管理员通讯录权限范围



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ege851

GET https://oapi.dingtalk.com/topapi/user/get_admin_scope?access_token=ACCESS_TOKEN
*/
func GetAdminScope(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetAdminScope)
}

/*
根据unionid获取userid



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ege851

GET https://oapi.dingtalk.com/user/getUseridByUnionid?access_token=ACCESS_TOKENunionid=xxx
*/
func GetUseridByUnionid(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetUseridByUnionid + "?" + params.Encode())
}

/*
根据手机号获取userid



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ege851

GET https://oapi.dingtalk.com/user/get_by_mobile?access_token=ACCESS_TOKENmobile=1xxxxxxxxxx
*/
func GetByMobile(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetByMobile + "?" + params.Encode())
}

/*
获取企业员工人数



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ege851

GET https://oapi.dingtalk.com/user/get_org_user_count?access_token=ACCESS_TOKENonlyActive=0
*/
func GetOrgUserCount(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetOrgUserCount + "?" + params.Encode())
}

/*
未登录钉钉的员工列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ege851

POST https://oapi.dingtalk.com/topapi/inactive/user/get?access_token=ACCESS_TOKEN
*/
func InactiveUserGet(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiInactiveUserGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}
