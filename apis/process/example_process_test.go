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

package process_test

import (
	"fmt"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/process"
)

func ExampleCreate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := process.Create(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleListIds() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := process.ListIds(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := process.Get(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetTodoNum() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := process.GetTodoNum(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleListByUserId() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := process.ListByUserId(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCspaceInfo() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := process.CspaceInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCspacePreview() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := process.CspacePreview(ctx, payload)

	fmt.Println(resp, err)
}
