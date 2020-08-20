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

package service_test

import (
	"fmt"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/private/service"
)

func ExampleServiceaccountAdd() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := service.ServiceaccountAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleServiceaccountGet() {
	var ctx *dingding.App

	resp, err := service.ServiceaccountGet(ctx)

	fmt.Println(resp, err)
}

func ExampleServiceaccountUpdate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := service.ServiceaccountUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleServiceaccountList() {
	var ctx *dingding.App

	resp, err := service.ServiceaccountList(ctx)

	fmt.Println(resp, err)
}

func ExampleMaterialArticleAdd() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := service.MaterialArticleAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMaterialArticleGet() {
	var ctx *dingding.App

	resp, err := service.MaterialArticleGet(ctx)

	fmt.Println(resp, err)
}

func ExampleMaterialArticleUpdate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := service.MaterialArticleUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMaterialArticlePublish() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := service.MaterialArticlePublish(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMaterialArticleList() {
	var ctx *dingding.App

	resp, err := service.MaterialArticleList(ctx)

	fmt.Println(resp, err)
}

func ExampleMaterialArticleDelete() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := service.MaterialArticleDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMaterialNewsAdd() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := service.MaterialNewsAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMaterialNewsGet() {
	var ctx *dingding.App

	resp, err := service.MaterialNewsGet(ctx)

	fmt.Println(resp, err)
}

func ExampleMaterialNewsUpdate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := service.MaterialNewsUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMaterialNewsList() {
	var ctx *dingding.App

	resp, err := service.MaterialNewsList(ctx)

	fmt.Println(resp, err)
}

func ExampleMaterialNewsDelete() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := service.MaterialNewsDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMessageMassSend() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := service.MessageMassSend(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMessageMassRecall() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := service.MessageMassRecall(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleServiceaccountMenuUpdate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := service.ServiceaccountMenuUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleServiceaccountMenuGet() {
	var ctx *dingding.App

	resp, err := service.ServiceaccountMenuGet(ctx)

	fmt.Println(resp, err)
}
