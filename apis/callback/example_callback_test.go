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

package callback_test

import (
	"fmt"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/callback"
)

func ExampleRegisterCallBack() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := callback.RegisterCallBack(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCallBack() {
	var ctx *dingding.App

	resp, err := callback.GetCallBack(ctx)

	fmt.Println(resp, err)
}

func ExampleUpdateCallBack() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := callback.UpdateCallBack(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDeleteCallBack() {
	var ctx *dingding.App

	resp, err := callback.DeleteCallBack(ctx)

	fmt.Println(resp, err)
}

func ExampleGetCallBackFailedResult() {
	var ctx *dingding.App

	resp, err := callback.GetCallBackFailedResult(ctx)

	fmt.Println(resp, err)
}
