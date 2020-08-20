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

// Package alitrip 阿里商旅
package alitrip

import (
	"bytes"

	"github.com/fastwego/dingding"
)

const (
	apiBtripFlightCitySuggest      = "/topapi/alitrip/btrip/flight/city/suggest"
	apiBtripTrainCitySuggest       = "/topapi/alitrip/btrip/train/city/suggest"
	apiBtripCostCenterNew          = "/topapi/alitrip/btrip/cost/center/new"
	apiABtripCostCenterModify      = "/topapi/alitrip/btrip/cost/center/modify"
	apiBtripCostCenterQuery        = "/topapi/alitrip/btrip/cost/center/query"
	apiBtripCostCenterDelete       = "/topapi/alitrip/btrip/cost/center/delete"
	apiBtripCostCenterEntitySet    = "/topapi/alitrip/btrip/cost/center/entity/set"
	apiBtripCostCenterEntityDelete = "/topapi/alitrip/btrip/cost/center/entity/delete"
	apiBtripCostCenterTransfer     = "/topapi/alitrip/btrip/cost/center/transfer"
)

/*
机票城市搜索



See: https://ding-doc.dingtalk.com/doc#/serverapi2/kg5g2g

POST https://oapi.dingtalk.com/topapi/alitrip/btrip/flight/city/suggest?access_token=ACCESS_TOKEN
*/
func BtripFlightCitySuggest(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBtripFlightCitySuggest, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
火车票城市搜索



See: https://ding-doc.dingtalk.com/doc#/serverapi2/kg5g2g

POST https://oapi.dingtalk.com/topapi/alitrip/btrip/train/city/suggest?access_token=ACCESS_TOKEN
*/
func BtripTrainCitySuggest(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBtripTrainCitySuggest, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
新建成本中心



See: https://ding-doc.dingtalk.com/doc#/serverapi2/kaklal

POST https://oapi.dingtalk.com/topapi/alitrip/btrip/cost/center/new?access_token=ACCESS_TOKEN
*/
func BtripCostCenterNew(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBtripCostCenterNew, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改成本中心



See: https://ding-doc.dingtalk.com/doc#/serverapi2/kaklal

POST https://oapi.dingtalk.com/topapi/alitrip/btrip/cost/center/modify?access_token=ACCESS_TOKEN
*/
func ABtripCostCenterModify(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiABtripCostCenterModify, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询成本中心



See: https://ding-doc.dingtalk.com/doc#/serverapi2/kaklal

POST https://oapi.dingtalk.com/topapi/alitrip/btrip/cost/center/query?access_token=ACCESS_TOKEN
*/
func BtripCostCenterQuery(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBtripCostCenterQuery, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除成本中心



See: https://ding-doc.dingtalk.com/doc#/serverapi2/kaklal

POST https://oapi.dingtalk.com/topapi/alitrip/btrip/cost/center/delete?access_token=ACCESS_TOKEN
*/
func BtripCostCenterDelete(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBtripCostCenterDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
设置成本中心人员信息



See: https://ding-doc.dingtalk.com/doc#/serverapi2/kaklal

POST https://oapi.dingtalk.com/topapi/alitrip/btrip/cost/center/entity/set?access_token=ACCESS_TOKEN
*/
func BtripCostCenterEntitySet(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBtripCostCenterEntitySet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除成本中心人员信息



See: https://ding-doc.dingtalk.com/doc#/serverapi2/kaklal

POST https://oapi.dingtalk.com/topapi/alitrip/btrip/cost/center/entity/delete?access_token=ACCESS_TOKEN
*/
func BtripCostCenterEntityDelete(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBtripCostCenterEntityDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
商旅成本中心转换为外部成本中心



See: https://ding-doc.dingtalk.com/doc#/serverapi2/kaklal

POST https://oapi.dingtalk.com/topapi/alitrip/btrip/cost/center/transfer?access_token=ACCESS_TOKEN
*/
func BtripCostCenterTransfer(ctx *dingding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBtripCostCenterTransfer, bytes.NewReader(payload), "application/json;charset=utf-8")
}
