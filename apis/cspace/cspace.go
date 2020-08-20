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

// Package cspace 文件存储
package cspace

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path"

	"github.com/fastwego/dingding"
)

const (
	apiUpload           = "/media/upload"
	apiAddToSingleChat  = "/cspace/add_to_single_chat"
	apiAdd              = "/cspace/add"
	apiGetCustomSpace   = "/cspace/get_custom_space"
	apiUsedInfo         = "/cspace/used_info"
	apiGrantCustomSpace = "/cspace/grant_custom_space"
	apiUploadSingle     = "/file/upload/single"
	apiTransaction      = "/file/upload/transaction"
	apiUploadChunk      = "/file/upload/chunk"
)

/*
上传媒体文件



See: https://ding-doc.dingtalk.com/doc#/serverapi2/bcmg0i

POST https://oapi.dingtalk.com/media/upload?access_token=ACCESS_TOKENtype=TYPE
*/
func Upload(ctx *dingding.App, media string, params url.Values) (resp []byte, err error) {

	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(media))
		if err != nil {
			return
		}
		file, err := os.Open(media)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

		// field
		err = m.WriteField("type", params.Get("type"))
		if err != nil {
			return
		}

	}()

	return ctx.Client.HTTPPost(apiUpload, r, m.FormDataContentType())

}

/*
发送钉盘文件给指定用户



See: https://ding-doc.dingtalk.com/doc#/serverapi2/wk3krc

POST https://oapi.dingtalk.com/cspace/add_to_single_chat?access_token=ACCESS_TOKENagent_id=AGENT_IDuserid=USERIDmedia_id=MEDIA_IDfile_name=FILE_NAME
*/
func AddToSingleChat(ctx *dingding.App, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddToSingleChat+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
新增文件到用户自定义空间



See: https://ding-doc.dingtalk.com/doc#/serverapi2/wk3krc

GET https://oapi.dingtalk.com/cspace/add?access_token=ACCESS_TOKENagent_id=AGENT_IDcode=CODEmedia_id=MEDIA_IDspace_id=SPACE_IDfolder_id=FOLDER_IDname=NAMEoverwrite=OVERWRITEpath=PATH
*/
func Add(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiAdd + "?" + params.Encode())
}

/*
获取企业下的自定义空间



See: https://ding-doc.dingtalk.com/doc#/serverapi2/wk3krc

GET https://oapi.dingtalk.com/cspace/get_custom_space?access_token=ACCESS_TOKENdomain=DOMAINagent_id=AGENT_ID
*/
func GetCustomSpace(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetCustomSpace + "?" + params.Encode())
}

/*
获取应用自定义空间使用详情



See: https://ding-doc.dingtalk.com/doc#/serverapi2/wk3krc

POST https://oapi.dingtalk.com/cspace/used_info?access_token=ACCESS_TOKENagent_id=AGENT_IDdomain=DOMAIN
*/
func UsedInfo(ctx *dingding.App, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUsedInfo+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
授权用户访问企业自定义空间



See: https://ding-doc.dingtalk.com/doc#/serverapi2/wk3krc

GET https://oapi.dingtalk.com/cspace/grant_custom_space?access_token=ACCESS_TOKENdomain=DOMAINagent_id=AGENT_IDtype=TYPEuserid=USERIDpath=PATHfileids=FILEIDSduration=DURATION
*/
func GrantCustomSpace(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGrantCustomSpace + "?" + params.Encode())
}

/*
单步上传文件



See: https://ding-doc.dingtalk.com/doc#/serverapi2/wk3krc

POST https://oapi.dingtalk.com/file/upload/single?access_token=ACCESS_TOKENagent_id=AGENT_IDfile_size=FILE_SIZE
*/
func UploadSingle(ctx *dingding.App, media string, params url.Values) (resp []byte, err error) {

	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(media))
		if err != nil {
			return
		}
		file, err := os.Open(media)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

		// field
		err = m.WriteField("agent_id", params.Get("agent_id"))
		if err != nil {
			return
		}

		err = m.WriteField("file_size", params.Get("file_size"))
		if err != nil {
			return
		}

	}()

	return ctx.Client.HTTPPost(apiUploadSingle, r, m.FormDataContentType())

}

/*
开启/提交文件上传事务



See: https://ding-doc.dingtalk.com/doc#/serverapi2/wk3krc

GET https://oapi.dingtalk.com/file/upload/transaction?access_token=ACCESS_TOKENagent_id=AGENT_IDfile_size=FILE_SIZEchunk_numbers=CHUNK_NUMBERS
*/
func Transaction(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiTransaction + "?" + params.Encode())
}

/*
上传文件块



See: https://ding-doc.dingtalk.com/doc#/serverapi2/wk3krc

POST https://oapi.dingtalk.com/file/upload/chunk?access_token=ACCESS_TOKENagent_id=AGENT_IDupload_id=UPLOAD_IDchunk_sequence=CHUNK_SEQUENCE
*/
func UploadChunk(ctx *dingding.App, media string, params url.Values) (resp []byte, err error) {

	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(media))
		if err != nil {
			return
		}
		file, err := os.Open(media)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

		// field
		err = m.WriteField("agent_id", params.Get("agent_id"))
		if err != nil {
			return
		}

		err = m.WriteField("upload_id", params.Get("upload_id"))
		if err != nil {
			return
		}

		err = m.WriteField("chunk_sequence", params.Get("chunk_sequence"))
		if err != nil {
			return
		}

	}()

	return ctx.Client.HTTPPost(apiUploadChunk, r, m.FormDataContentType())

}
