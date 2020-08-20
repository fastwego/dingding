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

// Package service 专属钉钉/互动服务窗
package service

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiServiceaccountAdd        = "/topapi/serviceaccount/add"
	apiServiceaccountGet        = "/topapi/serviceaccount/get"
	apiServiceaccountUpdate     = "/topapi/serviceaccount/update"
	apiServiceaccountList       = "/topapi/serviceaccount/list"
	apiMaterialArticleAdd       = "/topapi/material/article/add"
	apiMaterialArticleGet       = "/topapi/material/article/get"
	apiMaterialArticleUpdate    = "/topapi/material/article/update"
	apiMaterialArticlePublish   = "/topapi/material/article/publish"
	apiMaterialArticleList      = "/topapi/material/article/list"
	apiMaterialArticleDelete    = "/topapi/material/article/delete"
	apiMaterialNewsAdd          = "/topapi/material/news/add"
	apiMaterialNewsGet          = "/topapi/material/news/get"
	apiMaterialNewsUpdate       = "/topapi/material/news/update"
	apiMaterialNewsList         = "/topapi/material/news/list"
	apiMaterialNewsDelete       = "/topapi/material/news/delete"
	apiMessageMassSend          = "/topapi/message/mass/send"
	apiMessageMassRecall        = "/topapi/message/mass/recall"
	apiServiceaccountMenuUpdate = "/topapi/serviceaccount/menu/update"
	apiServiceaccountMenuGet    = "/topapi/serviceaccount/menu/get"
)

/*
新增服务号



See: https://ding-doc.dingtalk.com/doc#/serverapi2/elc75y

POST https://oapi.dingtalk.com/topapi/serviceaccount/add?access_token=ACCESS_TOKEN
*/
func ServiceaccountAdd(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiServiceaccountAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询服务号详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/elc75y

GET https://oapi.dingtalk.com/topapi/serviceaccount/get?access_token=ACCESS_TOKEN
*/
func ServiceaccountGet(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiServiceaccountGet)
}

/*
更新服务号



See: https://ding-doc.dingtalk.com/doc#/serverapi2/elc75y

POST https://oapi.dingtalk.com/topapi/serviceaccount/update?access_token=ACCESS_TOKEN
*/
func ServiceaccountUpdate(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiServiceaccountUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询服务号列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/elc75y

GET https://oapi.dingtalk.com/topapi/serviceaccount/list?access_token=ACCESS_TOKEN
*/
func ServiceaccountList(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiServiceaccountList)
}

/*
新增文章



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ngs3sp

POST https://oapi.dingtalk.com/topapi/material/article/add?access_token=ACCESS_TOKEN
*/
func MaterialArticleAdd(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMaterialArticleAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取文章详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ngs3sp

GET https://oapi.dingtalk.com/topapi/material/article/get?access_token=ACCESS_TOKEN
*/
func MaterialArticleGet(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiMaterialArticleGet)
}

/*
更新文章



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ngs3sp

POST https://oapi.dingtalk.com/topapi/material/article/update?access_token=ACCESS_TOKEN
*/
func MaterialArticleUpdate(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMaterialArticleUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
发布文章



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ngs3sp

POST https://oapi.dingtalk.com/topapi/material/article/publish?access_token=ACCESS_TOKEN
*/
func MaterialArticlePublish(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMaterialArticlePublish, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询文章列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ngs3sp

GET https://oapi.dingtalk.com/topapi/material/article/list?access_token=ACCESS_TOKEN
*/
func MaterialArticleList(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiMaterialArticleList)
}

/*
删除文章



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ngs3sp

POST https://oapi.dingtalk.com/topapi/material/article/delete?access_token=ACCESS_TOKEN
*/
func MaterialArticleDelete(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMaterialArticleDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
新增消息卡片



See: https://ding-doc.dingtalk.com/doc#/serverapi2/xwidee

POST https://oapi.dingtalk.com/topapi/material/news/add?access_token=ACCESS_TOKEN
*/
func MaterialNewsAdd(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMaterialNewsAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取消息卡片详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/xwidee

GET https://oapi.dingtalk.com/topapi/material/news/get?access_token=ACCESS_TOKEN
*/
func MaterialNewsGet(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiMaterialNewsGet)
}

/*
更新消息卡片



See: https://ding-doc.dingtalk.com/doc#/serverapi2/xwidee

POST https://oapi.dingtalk.com/topapi/material/news/update?access_token=ACCESS_TOKEN
*/
func MaterialNewsUpdate(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMaterialNewsUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询消息卡片列表



See: https://ding-doc.dingtalk.com/doc#/serverapi2/xwidee

GET https://oapi.dingtalk.com/topapi/material/news/list?access_token=ACCESS_TOKEN
*/
func MaterialNewsList(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiMaterialNewsList)
}

/*
删除消息卡片



See: https://ding-doc.dingtalk.com/doc#/serverapi2/xwidee

POST https://oapi.dingtalk.com/topapi/material/news/delete?access_token=ACCESS_TOKEN
*/
func MaterialNewsDelete(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMaterialNewsDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量发消息



See: https://ding-doc.dingtalk.com/doc#/serverapi2/aidmta

POST https://oapi.dingtalk.com/topapi/message/mass/send?access_token=ACCESS_TOKEN
*/
func MessageMassSend(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMessageMassSend, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
消息撤回



See: https://ding-doc.dingtalk.com/doc#/serverapi2/aidmta

POST https://oapi.dingtalk.com/topapi/message/mass/recall?access_token=ACCESS_TOKEN
*/
func MessageMassRecall(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMessageMassRecall, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新菜单



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ux6nqr

POST https://oapi.dingtalk.com/topapi/serviceaccount/menu/update?access_token=ACCESS_TOKEN
*/
func ServiceaccountMenuUpdate(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiServiceaccountMenuUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取菜单



See: https://ding-doc.dingtalk.com/doc#/serverapi2/ux6nqr

GET https://oapi.dingtalk.com/topapi/serviceaccount/menu/get?access_token=ACCESS_TOKEN
*/
func ServiceaccountMenuGet(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiServiceaccountMenuGet)
}
