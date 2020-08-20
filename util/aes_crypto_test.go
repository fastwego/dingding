// Copyright 2014 chanxuehong(chanxuehong@gmail.com)
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

package util

import (
	"bytes"
	"fmt"
	"testing"
)

func TestAESEncryptMsg(t *testing.T) {
	appId := "wx45f133bf6fce646e"
	encodingAESKey := "AdiqDDDvUNCeE1ZW5XJmjf9fqNBJpGBs4vL4cHKmHBS"

	random := []byte("cc9632a98304f81c")
	plaintext := []byte(`<xml><ToUserName><![CDATA[gh_b1eb3f8bd6c6]]></ToUserName>
<FromUserName><![CDATA[okPEat9FRX96xG8JQvTxHLpzDV64]]></FromUserName>
<CreateTime>1458889120</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[test text message]]></Content>
<MsgId>6265881059295447969</MsgId>
</xml>`)

	wantBase64Ciphertext := "gsKjKVChgrDDifoOdfrL/MWujrKxnSR1jBopv0zEdHFQXcx5I0bf4UxIRjektLHEziRxpU5zTw0s+gYF3WOFjk" +
		"In30gzQl9XNKr2A+DCHVG05I2EcbOnnAR5EYgjGLAipra5nOfCPPIRFQTZ7SdanniXX73YOiCAKJlNH21+PApcYu4rPXxJ4eTXbBLwmfnS" +
		"l7iojDX1LQIcC3FYmaapMQq/u+sJGsxshp4dLXJ6A5Ji3cSYAzXRVIxbNlHN1MWfdcZ0O5+ZtOU6dZiD8hJ4kkxh05EfOedWFjUy7ZhmXg" +
		"rpZ4WpYPsrythXHE2Bg1Ohz8uf5h5X31yuU/3FPoa8rD21pfZnAjBT1QCkn6MxtL5lR+yoQReLwElVbtB6yJFPIZ+n9Qh/yKfIasxkzgIE" +
		"o0pwPEbS17WCTyvRItJtU6tlo3rawJX+fsV63tKIfmLZjyZPDlV9/ka/nA/DT0KODw=="

	base64Ciphertext := AESEncryptMsg(random, plaintext, appId, encodingAESKey)
	if base64Ciphertext != wantBase64Ciphertext {
		t.Errorf("tests AESEncryptMsg failed,\nhave: %s\nwant: %s\n", base64Ciphertext, wantBase64Ciphertext)
		return
	}
}

func TestAESDecryptMsg(t *testing.T) {
	encodingAesKey := "4g5j64qlyl3zvetqxz5jiocdr586fn2zvjpa8zls3ij"

	wantRandom := []byte("hU3bEfGZZewzhG5a")
	wantPlaintext := []byte(`{"EventType":"check_create_suite_url","Random":"LPIdSnlF","TestSuiteKey":"suite4xxxxxxxxxxxxxxx"}`)
	wantAppId := []byte("suite4xxxxxxxxxxxxxxx")

	base64Ciphertext := "1a3NBxmCFwkCJvfoQ7WhJHB+iX3qHPsc9JbaDznE1i03peOk1LaOQoRz3+nlyGNhwmwJ3vDMG+OzrHMeiZI7gTRWVdUBmfxjZ8Ej23JVYa9VrYeJ5as7XM/ZpulX8NEQis44w53h1qAgnC3PRzM7Zc/D6Ibr0rgUathB6zRHP8PYrfgnNOS9PhSBdHlegK+AGGanfwjXuQ9+0pZcy0w9lQ=="
	random, plaintext, appId, err := AESDecryptMsg(base64Ciphertext, encodingAesKey)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(random, wantRandom) {
		t.Errorf("tests AESDecryptMsg failed,\nhave random: %s\nwant random: %s\n", random, wantRandom)
	}
	if !bytes.Equal(plaintext, wantPlaintext) {
		fmt.Println(len(plaintext), " == ", len(wantPlaintext))
		t.Errorf("tests AESDecryptMsg failed,\nhave plaintext: %s\nwant plaintext: %s\n", plaintext, wantPlaintext)
	}
	if !bytes.Equal(appId, wantAppId) {
		t.Errorf("tests AESDecryptMsg failed,\nhave appid: %s\nwant appid: %s\n", appId, wantAppId)
	}
}
