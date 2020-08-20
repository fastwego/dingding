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

package role_test

import (
	"fmt"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/contact/role"
)

func ExampleList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := role.List(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSimpleList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := role.SimpleList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetRoleGroup() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := role.GetRoleGroup(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetRole() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := role.GetRole(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddRole() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := role.AddRole(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateRole() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := role.UpdateRole(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDeleteRole() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := role.DeleteRole(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddRoleGroup() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := role.AddRoleGroup(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddRolesForEmps() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := role.AddRolesForEmps(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleRemoveRolesForEmps() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := role.RemoveRolesForEmps(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleScopeUpdate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := role.ScopeUpdate(ctx, payload)

	fmt.Println(resp, err)
}
