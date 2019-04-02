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
)

type DeleteTopicRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /*   */
    Name string `json:"name"`
}

/*
 * param regionId: 地域ID (Required)
 * param name:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteTopicRequest(
    regionId string,
    name string,
) *DeleteTopicRequest {

	return &DeleteTopicRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/topic",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param name:  (Required)
 */
func NewDeleteTopicRequestWithAllParams(
    regionId string,
    name string,
) *DeleteTopicRequest {

    return &DeleteTopicRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topic",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteTopicRequestWithoutParam() *DeleteTopicRequest {

    return &DeleteTopicRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topic",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DeleteTopicRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param name: (Required) */
func (r *DeleteTopicRequest) SetName(name string) {
    r.Name = name
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteTopicRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteTopicResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteTopicResult `json:"result"`
}

type DeleteTopicResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
}