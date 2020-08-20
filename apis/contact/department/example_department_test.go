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

package department_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/contact/department"
)

func ExampleCreate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := department.Create(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := department.Update(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := department.Delete(ctx, params)

	fmt.Println(resp, err)
}

func ExampleListIds() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := department.ListIds(ctx, params)

	fmt.Println(resp, err)
}

func ExampleList() {
	var ctx *dingding.App

	resp, err := department.List(ctx)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := department.Get(ctx, params)

	fmt.Println(resp, err)
}

func ExampleListParentDeptsByDept() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := department.ListParentDeptsByDept(ctx, params)

	fmt.Println(resp, err)
}

func ExampleListParentDepts() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := department.ListParentDepts(ctx, params)

	fmt.Println(resp, err)
}
