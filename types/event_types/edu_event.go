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

const (
	EventTypeEduUserInsert = "edu_user_insert" // 人员身份新增

	EventTypeEduUserUpdate = "edu_user_update" //人员身份更新

	EventTypeEduUserDelete = "edu_user_delete" //人员身份删除

	EventTypeEduUserRelationInsert = "edu_user_relation_insert" //人员关系新增

	EventTypeEduUserRelationUpdate = "edu_user_relation_update" //人员关系更新

	EventTypeEduUserRelationDelete = "edu_user_relation_delete" //人员关系删除

	EventTypeEduDeptInsert = "edu_dept_insert" //部门节点新增

	EventTypeEduDeptUpdate = "edu_dept_update" //部门节点更新

	EventTypeEduDeptDelete = "edu_dept_delete" // 部门节点删除

)
