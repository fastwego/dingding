// Copyright 2021 FastWeGo
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

package dingding

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	contentTypeApplicationJSON = "application/json"
)

// ServerURL DingDing api服务地址
var ServerURL = "https://oapi.dingtalk.com"

// UserAgent 发送请求是，使用的用户标识
var UserAgent = "fastwego/dingding"

var errorSystemBusy = errors.New("system busy")

// NewClient 声明一个客户端
func NewClient(AccessTokenManager AccessTokenManager) (client *Client) {
	return &Client{
		AccessTokenManager: AccessTokenManager,
		HTTPClient:         http.DefaultClient,
	}
}

// Client 用于向接口发送请求
type Client struct {
	AccessTokenManager AccessTokenManager
	HTTPClient         *http.Client
}

// Do 执行 请求
func (client *Client) Do(req *http.Request) (resp []byte, err error) {

	// 添加 serverUrl
	if !strings.HasPrefix(req.URL.String(), "http") {
		parse, _ := url.Parse(ServerURL)
		req.URL.Host = parse.Host
		req.URL.Scheme = parse.Scheme
	}

	// 默认 Header Content-Type
	if req.Method == http.MethodPost && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", contentTypeApplicationJSON)
	}

	// 添加 access_token
	accessToken, err := client.AccessTokenManager.GetAccessToken()
	if err != nil {
		return
	}
	q := req.URL.Query()
	q.Set(client.AccessTokenManager.GetName(), accessToken)
	req.URL.RawQuery = q.Encode()

	// 添加 User-Agent
	req.Header.Add("User-Agent", UserAgent)

	if Logger != nil {
		Logger.Printf("%s %s %v", req.Method, req.URL.String(), req.Header)
	}

	response, err := client.HTTPClient.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	resp, err = responseFilter(response)

	// -1 系统繁忙，此时请开发者稍候再试
	// 重试一次
	if err == errorSystemBusy {

		if Logger != nil {
			Logger.Printf("%v : retry %s %s Headers %v", errorSystemBusy, req.Method, req.URL.String(), req.Header)
		}

		response, err = client.HTTPClient.Do(req)
		if err != nil {
			return
		}

		resp, err = responseFilter(response)
	}

	return
}

/*
筛查钉钉 api 服务器响应，判断以下错误：

- http 状态码 不为 200

- 接口响应错误码 errcode 不为 0
*/
func responseFilter(response *http.Response) (resp []byte, err error) {

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("response.Status %s", response.Status)
		return
	}

	//log.Println(response.Header)

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if response.Header.Get("Content-Type") == "application/json" { // 只针对 json
		errorResponse := struct {
			Errcode int64  `json:"errcode"`
			Errmsg  string `json:"errmsg"`
		}{}
		err = json.Unmarshal(resp, &errorResponse)
		if err != nil {
			return
		}

		//  -1	系统繁忙，此时请开发者稍候再试
		if errorResponse.Errcode == -1 {
			err = errorSystemBusy
			return
		}

		if errorResponse.Errcode != 0 {
			err = errors.New(string(resp))
			return
		}
	}
	return
}
