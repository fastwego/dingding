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

package datav_test

import (
	"fmt"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/private/datav"
)

func ExampleGetAccessToken() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.GetAccessToken(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacDatavDeptVideoliveList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacDatavDeptVideoliveList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacDatavVideoliveDetailList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacDatavVideoliveDetailList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacDatavVideoliveViewerList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacDatavVideoliveViewerList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacDatavGroupGet() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacDatavGroupGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacV2DatavVideoconfGet() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacV2DatavVideoconfGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacDatavDeptVideoconfList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacDatavDeptVideoconfList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacDatavVideoconfDetailList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacDatavVideoconfDetailList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacDatavTelconfGet() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacDatavTelconfGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacDatavDeptTelconfList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacDatavDeptTelconfList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacDatavTelconfDetailList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacDatavTelconfDetailList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacDatavChatSummaryGet() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacDatavChatSummaryGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacDatavMicroappDetailList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacDatavMicroappDetailList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacDatavDauSummaryGet() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacDatavDauSummaryGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacDatavDeptDauList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacDatavDeptDauList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKacDatavInactivatedUserList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := datav.KacDatavInactivatedUserList(ctx, payload)

	fmt.Println(resp, err)
}
