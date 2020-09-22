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

package message_test

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/message"
)

func ExampleAsyncsendV2() {
	var ctx *dingding.App

	type Msg struct {
		Msgtype    string `json:"msgtype"`
		ActionCard struct {
			Title       string `json:"title"`
			Markdown    string `json:"markdown"`
			SingleTitle string `json:"single_title"`
			SingleURL   string `json:"single_url"`
		} `json:"action_card"`
	}
	msg := Msg{}
	msg.Msgtype = "action_card"
	msg.ActionCard.Title = "Title"
	msg.ActionCard.Markdown = `# 今天晚上不见不散`
	msg.ActionCard.SingleTitle = "马上看看"
	msg.ActionCard.SingleURL = "https://fastwego.dev"

	data := struct {
		AgentId    string `json:"agent_id"`
		UseridList string `json:"userid_list"`
		Msg        Msg    `json:"msg"`
	}{
		AgentId:    "AGENT_ID",
		UseridList: "USERID",
	}
	data.Msg = msg
	payload, err := json.Marshal(data)
	resp, err := message.AsyncsendV2(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetSendProgress() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := message.GetSendProgress(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetSendResult() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := message.GetSendResult(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleRecall() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := message.Recall(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSend() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := message.Send(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetReadList() {
	var ctx *dingding.App

	params := url.Values{}
	resp, err := message.GetReadList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleSendToConversation() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := message.SendToConversation(ctx, payload)

	fmt.Println(resp, err)
}
