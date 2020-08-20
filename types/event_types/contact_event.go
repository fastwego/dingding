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

package event_types

type EventContact struct {
	Event
	TimeStamp int64    `json:"TimeStamp"`
	UserID    []string `json:"UserId"`
	DeptId    []string `json:"DeptId"`
	CorpID    string   `json:"CorpId"`
}

const (
	EventTypeUserAddOrg = "user_add_org" //用户变更 -通讯录用户增加

	EventTypeUserModifyOrg = "user_modify_org" // 用户变更 - 通讯录用户更改

	EventTypeUserLeaveOrg = "user_leave_org" //用户变更 - 通讯录用户离职

	EventTypeUserActiveOrg = "user_active_org" //用户变更 - 加入企业后用户激活

	EventTypeOrgAdminAdd = "org_admin_add" //用户变更 - 通讯录用户被设为管理员

	EventTypeOrgAdminRemoce = "org_admin_remove" // 用户变更 - 通讯录用户被取消设置管理员

	EventTypeOrgDeptCreate = "org_dept_create" // 部门变更 - 通讯录企业部门创建

	EventTypeOrgDeptModify = "org_dept_modify" //部门变更 - 通讯录企业部门修改

	EventTypeOrgDeptRemove = "org_dept_remove" // 部门变更 - 通讯录企业部门删除

	EventTypeOrgRemove = "org_remove" // 企业信息变更 - 企业被解散

	EventTypeOrgChange = "org_change" // 企业信息变更 - 企业信息发生变更

	EventTypeLeaveUserChange = "label_user_change" //角色变更 - 员工角色信息发生变更

	EventTypeLabelConfAdd = "label_conf_add" //角色变更 - 增加角色或者角色组

	EventTypeLabelConfDel = "label_conf_del" //角色变更 - 删除角色或者角色组

	EventTypeLabelConfModify = "label_conf_modify" //角色变更 - 修改角色或者角色组
)
