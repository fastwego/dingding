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
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fastwego/dingding/util"
)

// Crypto 加密结构体
type Crypto struct {
	Token          string
	EncodingAESKey string
	SuiteKey       string
}

// NewCrypto 创建加解密处理器
func NewCrypto(token, encodingAESKey, suiteKey string) *Crypto {
	return &Crypto{
		Token:          token,
		EncodingAESKey: encodingAESKey,
		SuiteKey:       suiteKey,
	}
}

// GetEncryptMsg 加密消息
func (c *Crypto) GetEncryptMsg(msg string) (resp map[string]string) {

	// msg_signature=sha1(sort(token、timestamp、nonce、msg_encrypt))

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := util.GetRandString(16)

	encryptMsg := util.AESEncryptMsg([]byte(nonce), []byte(msg), c.SuiteKey, c.EncodingAESKey)

	/*
		返回给钉钉的数据说明：
		{
		  "msg_signature":"111108bb8e6dbce3c9671d6fdb69d15066227608",
		  "timeStamp":"1783610513",
		  "nonce":"123456",
		  "encrypt":"1ojQf0NSvw2WPvfghjkl..."
		  }
	*/
	return map[string]string{
		"msg_signature": c.generateSignature(timestamp, nonce, encryptMsg),
		"timeStamp":     timestamp, // !!! 钉钉的 bug S 要求大写 否则钉钉服务器校验不通过
		"nonce":         nonce,
		"encrypt":       encryptMsg,
	}
}

func (c *Crypto) generateSignature(timestamp, nonce, encryptMsg string) (signature string) {

	strs := []string{timestamp, nonce, c.Token, encryptMsg}
	sort.Strings(strs)

	h := sha1.New()
	_, _ = io.WriteString(h, strings.Join(strs, ""))
	signature = fmt.Sprintf("%x", h.Sum(nil))
	return
}

// GetDecryptMsg 解密消息
func (c *Crypto) GetDecryptMsg(timestamp, nonce, signature, encryptMsg string) (decryptMsg []byte, err error) {

	// 校验签名
	if c.generateSignature(timestamp, nonce, encryptMsg) != signature {
		return nil, fmt.Errorf("signature fail")
	}

	// 解密
	_, decryptMsg, _, err = util.AESDecryptMsg(encryptMsg, c.EncodingAESKey)
	if err != nil {
		return
	}

	if Logger != nil {
		Logger.Printf("AESDecryptMsg %s => %s", encryptMsg, string(decryptMsg))
	}

	return
}
