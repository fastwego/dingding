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

// Package contact 通讯录管理
package contact

import "github.com/fastwego/dingding"

const (
	apiScopes = "/auth/scopes"
)

/*
获取通讯录权限范围



See: https://ding-doc.dingtalk.com/doc#/serverapi2/vt6v7m

GET https://oapi.dingtalk.com/auth/scopes?access_token=ACCESS_TOKEN
*/
func Scopes(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiScopes)
}
