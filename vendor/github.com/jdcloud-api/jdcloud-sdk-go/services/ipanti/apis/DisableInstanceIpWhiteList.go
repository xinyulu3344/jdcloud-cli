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

type DisableInstanceIpWhiteListRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 实例id  */
    InstanceId string `json:"instanceId"`
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: 实例id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDisableInstanceIpWhiteListRequest(
    regionId string,
    instanceId string,
) *DisableInstanceIpWhiteListRequest {

	return &DisableInstanceIpWhiteListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:disableIpWhiteList",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: 实例id (Required)
 */
func NewDisableInstanceIpWhiteListRequestWithAllParams(
    regionId string,
    instanceId string,
) *DisableInstanceIpWhiteListRequest {

    return &DisableInstanceIpWhiteListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:disableIpWhiteList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDisableInstanceIpWhiteListRequestWithoutParam() *DisableInstanceIpWhiteListRequest {

    return &DisableInstanceIpWhiteListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:disableIpWhiteList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DisableInstanceIpWhiteListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例id(Required) */
func (r *DisableInstanceIpWhiteListRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisableInstanceIpWhiteListRequest) GetRegionId() string {
    return r.RegionId
}

type DisableInstanceIpWhiteListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisableInstanceIpWhiteListResult `json:"result"`
}

type DisableInstanceIpWhiteListResult struct {
}