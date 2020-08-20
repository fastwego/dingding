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

package user_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/contact/user"
)

func ExampleCreate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := user.Create(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := user.Update(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := user.Delete(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := user.Get(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetDeptMember() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := user.GetDeptMember(ctx, params)

	fmt.Println(resp, err)
}

func ExampleSimpleList() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := user.SimpleList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleListByPage() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := user.ListByPage(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetAdmin() {
	var ctx *dingding.App

	resp, err := user.GetAdmin(ctx)

	fmt.Println(resp, err)
}

func ExampleGetAdminScope() {
	var ctx *dingding.App

	resp, err := user.GetAdminScope(ctx)

	fmt.Println(resp, err)
}

func ExampleGetUseridByUnionid() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := user.GetUseridByUnionid(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetByMobile() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := user.GetByMobile(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetOrgUserCount() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := user.GetOrgUserCount(ctx, params)

	fmt.Println(resp, err)
}

func ExampleInactiveUserGet() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := user.InactiveUserGet(ctx, payload)

	fmt.Println(resp, err)
}
