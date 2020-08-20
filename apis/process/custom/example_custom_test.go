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

package custom_test

import (
	"fmt"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/process/custom"
)

func ExampleSave() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := custom.Save(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetByName() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := custom.GetByName(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := custom.Delete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCreate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := custom.Create(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := custom.Update(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBatchUpdate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := custom.BatchUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTaskCreate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := custom.TaskCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTaskUpdate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := custom.TaskUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTaskGroupCancel() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := custom.TaskGroupCancel(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTaskQuery() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := custom.TaskQuery(ctx, payload)

	fmt.Println(resp, err)
}
