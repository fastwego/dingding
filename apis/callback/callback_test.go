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

package callback

import (
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/test"
)

func TestMain(m *testing.M) {
	test.Setup()
	os.Exit(m.Run())
}

func TestRegisterCallBack(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiRegisterCallBack, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		ctx     *dingding.App
		payload []byte
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := RegisterCallBack(tt.args.ctx, tt.args.payload)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterCallBack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("RegisterCallBack() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestGetCallBack(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiGetCallBack, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		ctx *dingding.App
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := GetCallBack(tt.args.ctx)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCallBack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetCallBack() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestUpdateCallBack(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiUpdateCallBack, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		ctx     *dingding.App
		payload []byte
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := UpdateCallBack(tt.args.ctx, tt.args.payload)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCallBack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UpdateCallBack() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestDeleteCallBack(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiDeleteCallBack, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		ctx *dingding.App
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := DeleteCallBack(tt.args.ctx)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteCallBack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("DeleteCallBack() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestGetCallBackFailedResult(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiGetCallBackFailedResult, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		ctx *dingding.App
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := GetCallBackFailedResult(tt.args.ctx)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCallBackFailedResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetCallBackFailedResult() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
