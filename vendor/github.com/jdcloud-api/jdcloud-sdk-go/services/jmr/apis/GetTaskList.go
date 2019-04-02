// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    jmr "github.com/jdcloud-api/jdcloud-sdk-go/services/jmr/models"
)

type GetTaskListRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 需要查询的jobId  */
    JobId string `json:"jobId"`

    /*  (Optional) */
    SelectParams *jmr.SelectParams `json:""`
}

/*
 * param regionId: 地域ID (Required)
 * param jobId: 需要查询的jobId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetTaskListRequest(
    regionId string,
    jobId string,
) *GetTaskListRequest {

	return &GetTaskListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/{jobId}/task:list",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        JobId: jobId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param jobId: 需要查询的jobId (Required)
 * param selectParams:  (Optional)
 */
func NewGetTaskListRequestWithAllParams(
    regionId string,
    jobId string,
    selectParams *jmr.SelectParams,
) *GetTaskListRequest {

    return &GetTaskListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/{jobId}/task:list",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        JobId: jobId,
        SelectParams: selectParams,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetTaskListRequestWithoutParam() *GetTaskListRequest {

    return &GetTaskListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/{jobId}/task:list",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetTaskListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param jobId: 需要查询的jobId(Required) */
func (r *GetTaskListRequest) SetJobId(jobId string) {
    r.JobId = jobId
}

/* param selectParams: (Optional) */
func (r *GetTaskListRequest) SetSelectParams(selectParams *jmr.SelectParams) {
    r.SelectParams = selectParams
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetTaskListRequest) GetRegionId() string {
    return r.RegionId
}

type GetTaskListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetTaskListResult `json:"result"`
}

type GetTaskListResult struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Data jmr.TaskViewListData `json:"data"`
}