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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"

	"github.com/iancoleman/strcase"
)

var apiUniqMap = map[string]bool{}

func run() {
	docs := struct {
		Success   bool        `json:"success"`
		ErrorCode interface{} `json:"errorCode"`
		ErrorMsg  interface{} `json:"errorMsg"`
		Result    []struct {
			Title string `json:"title"`
			Slug  string `json:"slug"`
		} `json:"result"`
		Arguments interface{} `json:"arguments"`
	}{}

	resp, err := http.Get("https://ding-doc.dingtalk.com/getLeftNav.json?namespace=serverapi2")
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &docs)
	if err != nil {
		fmt.Println(string(body))
		panic(err)
	}

	for _, doc := range docs.Result {
		//fmt.Println(match[1])

		if doc.Slug == "#" {
			continue
		}

		link := "https://ding-doc.oss-cn-beijing.aliyuncs.com/0.0.368/bgrkdl/serverapi2_" + doc.Slug + ".html?_=1597895345802"
		resp, err := http.Get(link)
		if err != nil {
			panic(err)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
			continue
		}

		reg := regexp.MustCompile(`(?sU)(<h[1-3][^>]*>(.+)</h[1-3]>.+)?请求方式.+(GET|POST).+请求地址.+(https://oapi.dingtalk.com/\S+)<`)
		apiMatched := reg.FindAllStringSubmatch(string(body), -1)
		if apiMatched == nil {
			fmt.Println("match failed")
			continue
		}

		for _, api := range apiMatched {
			if len(api) < 5 {
				fmt.Println("len(api) < 5")
				continue
			}

			//fmt.Println(api[0])

			uri := strings.ReplaceAll(api[4], "&amp;", "&")
			parse, _ := url.Parse(uri)
			if _, ok := apiUniqMap[parse.Path]; !ok {
				apiUniqMap[parse.Path] = true
			} else {
				continue
			}
			_NAME_ := doc.Title
			if api[2] != "" && api[2] != "概述" && api[2] != "权限申请" {
				_NAME_ = strip.StripTags(api[2])
			}
			_DESCRIPTION_ := ""

			_REQUEST_ := api[3] + " " + uri
			_SEE_ := "https://ding-doc.dingtalk.com/doc#/serverapi2/" + doc.Slug
			_FUNCNAME_ := strcase.ToCamel(path.Base(parse.Path))

			_GET_PARAMS_ := ""
			uniqParams := map[string]bool{}
			for param_name, _ := range parse.Query() {
				if uniqParams[param_name] || strings.ToLower(param_name) == "access_token" {
					continue
				}
				_GET_PARAMS_ += "\t\t{Name: `" + param_name + "`, Type: `string`},\n"

				uniqParams[param_name] = true
			}
			if _GET_PARAMS_ != "" {
				_GET_PARAMS_ = `GetParams: []Param{
` + _GET_PARAMS_ + "\t},"
			}

			tpl := strings.ReplaceAll(itemTpl, "_NAME_", _NAME_)
			tpl = strings.ReplaceAll(tpl, "_DESCRIPTION_", _DESCRIPTION_)
			tpl = strings.ReplaceAll(tpl, "_REQUEST_", _REQUEST_)
			tpl = strings.ReplaceAll(tpl, "_SEE_", _SEE_)
			tpl = strings.ReplaceAll(tpl, "_FUNCNAME_", _FUNCNAME_)
			tpl = strings.ReplaceAll(tpl, "_GET_PARAMS_", _GET_PARAMS_)

			fmt.Println(tpl)
		}
	}

}

var itemTpl = `{
	Name:        "_NAME_",
	Description: "_DESCRIPTION_",
	Request:     "_REQUEST_",
	See:         "_SEE_",
	FuncName:    "_FUNCNAME_",
	_GET_PARAMS_
},`
