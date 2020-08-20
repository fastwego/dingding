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

package hr_test

import (
	"fmt"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/hr"
)

func ExampleEmployeeList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := hr.EmployeeList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleEmployeeUpdate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := hr.EmployeeUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleEmployeeFieldGrouplist() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := hr.EmployeeFieldGrouplist(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleQueryPreentry() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := hr.QueryPreentry(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleQueryonjob() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := hr.Queryonjob(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleQueryDimission() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := hr.QueryDimission(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleListDimission() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := hr.ListDimission(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddPreentry() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := hr.AddPreentry(ctx, payload)

	fmt.Println(resp, err)
}
