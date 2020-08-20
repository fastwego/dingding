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

package cspace_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/cspace"
)

func ExampleAddToSingleChat() {
	var ctx *dingding.App

	payload := []byte("{}")
	params := url.Values{}
	resp, err := cspace.AddToSingleChat(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleAdd() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := cspace.Add(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetCustomSpace() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := cspace.GetCustomSpace(ctx, params)

	fmt.Println(resp, err)
}

func ExampleUsedInfo() {
	var ctx *dingding.App

	payload := []byte("{}")
	params := url.Values{}
	resp, err := cspace.UsedInfo(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleGrantCustomSpace() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := cspace.GrantCustomSpace(ctx, params)

	fmt.Println(resp, err)
}

func ExampleTransaction() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := cspace.Transaction(ctx, params)

	fmt.Println(resp, err)
}
