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

package attendance_test

import (
	"fmt"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/attendance"
)

func ExampleListSchedule() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.ListSchedule(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleListByDay() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.ListByDay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleListByUsers() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.ListByUsers(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleScheduleAsync() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.ScheduleAsync(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleListByIds() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.ListByIds(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAttendanceShiftList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.AttendanceShiftList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAttendanceShiftQuery() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.AttendanceShiftQuery(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAttendanceShiftSearch() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.AttendanceShiftSearch(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleListRecord() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.ListRecord(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAttendanceList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.AttendanceList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetLeaveStatus() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.GetLeaveStatus(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetSimpleGroups() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.GetSimpleGroups(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAttendanceGroupMinimalismList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.AttendanceGroupMinimalismList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAttendanceGroupSearch() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.AttendanceGroupSearch(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAttendanceGroupQuery() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.AttendanceGroupQuery(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUserGroup() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.GetUserGroup(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAttendanceGroupMemberUsersList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.AttendanceGroupMemberUsersList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAttendanceGroupMemberListbyids() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.AttendanceGroupMemberListbyids(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAttendanceGroupMemberUpdate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.AttendanceGroupMemberUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAttendanceGroupMemberList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.AttendanceGroupMemberList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleIsOpenSmartReport() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.IsOpenSmartReport(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetAttColumns() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.GetAttColumns(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetColumnVal() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.GetColumnVal(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetLeaveTimeByNames() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := attendance.GetLeaveTimeByNames(ctx, payload)

	fmt.Println(resp, err)
}
