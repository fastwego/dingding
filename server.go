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

package dingding

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fastwego/dingding/types/event_types"
	"github.com/fastwego/dingding/util"
)

/*
响应钉钉请求 或 推送消息/事件 的服务器
*/
type Server struct {
	Ctx *App
}

// CheckUrl 服务器接口校验
func (s *Server) CheckUrl(writer http.ResponseWriter, request *http.Request) {

	// dev_msg_signature=sha1(sort(token、timestamp、nonce、msg_encrypt))

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := util.GetRandString(16)

	msgEncrypt := util.AESEncryptMsg([]byte(nonce), []byte("success"), s.Ctx.Config.CorpId, s.Ctx.Config.EncodingAESKey)

	strs := []string{timestamp, nonce, s.Ctx.Config.Token, msgEncrypt}
	sort.Strings(strs)

	h := sha1.New()
	_, _ = io.WriteString(h, strings.Join(strs, ""))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	/*
		返回给钉钉的数据说明：

		{
		  "msg_signature":"111108bb8e6dbce3c9671d6fdb69d15066227608",
		  "timeStamp":"1783610513",
		  "nonce":"123456",
		  "encrypt":"1ojQf0NSvw2WPvfghjkl..."
		  }
	*/
	echo := struct {
		MsgSignature string `json:"msg_signature"`
		TimeStamp    string `json:"timeStamp"`
		Nonce        string `json:"nonce"`
		Encrypt      string `json:"encrypt"`
	}{
		MsgSignature: signature,
		TimeStamp:    timestamp,
		Nonce:        nonce,
		Encrypt:      msgEncrypt,
	}

	data, err := json.Marshal(echo)
	if err != nil {
		return
	}
	if s.Ctx.Logger != nil {
		s.Ctx.Logger.Println(string(data))
	}

	writer.Write(data)
}

// ParseEvent 解析钉钉推送过来的消息/事件
func (s *Server) ParseEvent(body []byte) (eventType string, err error) {

	if s.Ctx.Logger != nil {
		s.Ctx.Logger.Println(string(body))
	}

	// 加密消息
	encryptMsg := struct {
		Encrypt string `json:"encrypt"`
	}{}
	err = json.Unmarshal(body, &encryptMsg)
	if err != nil {
		return
	}

	// 解密
	if encryptMsg.Encrypt != "" {
		_, body, _, err = util.AESDecryptMsg(encryptMsg.Encrypt, s.Ctx.Config.EncodingAESKey)
		if err != nil {
			s.Ctx.Logger.Println(string(body), err)
			return
		}

		if s.Ctx.Logger != nil {
			s.Ctx.Logger.Println("AESDecryptMsg ", encryptMsg.Encrypt, string(body))
		}

	}

	event := event_types.Event{}
	err = json.Unmarshal(body, &event)

	if err != nil {
		return
	}

	return event.EventType, nil
}
