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

package event_types

type EventChat struct {
	Event
	TimeStamp int64    `json:"TimeStamp"`
	CorpID    string   `json:"CorpId"`
	ChatID    string   `json:"ChatId"`
	UserID    []string `json:"UserId"`
	Operator  string   `json:"Operator"`
	Owner     string   `json:"Owner"`
	Title     string   `json:"Title"`
	agentId   string   `json:"agentId"`
}

const (
	EventTypeChatAddMember = "chat_add_member" //群会话添加人员

	EventTypeChatRemoveMember = "chat_remove_member" //群会话删除人员

	EventTypeChatQuit = "chat_quit" //群会话用户主动退群

	EventTypeChatUpdateOwner = "chat_update_owner" //群会话更换群主

	EventTypeChatUpdateTitle = "chat_update_title" //群会话更换群名称

	EventTypeChatDisband = "chat_disband" //群会话解散群
)
