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

package alitrip_test

import (
	"fmt"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/alitrip"
)

func ExampleBtripFlightCitySuggest() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := alitrip.BtripFlightCitySuggest(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBtripTrainCitySuggest() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := alitrip.BtripTrainCitySuggest(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBtripCostCenterNew() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := alitrip.BtripCostCenterNew(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleABtripCostCenterModify() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := alitrip.ABtripCostCenterModify(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBtripCostCenterQuery() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := alitrip.BtripCostCenterQuery(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBtripCostCenterDelete() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := alitrip.BtripCostCenterDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBtripCostCenterEntitySet() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := alitrip.BtripCostCenterEntitySet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBtripCostCenterEntityDelete() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := alitrip.BtripCostCenterEntityDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBtripCostCenterTransfer() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := alitrip.BtripCostCenterTransfer(ctx, payload)

	fmt.Println(resp, err)
}
