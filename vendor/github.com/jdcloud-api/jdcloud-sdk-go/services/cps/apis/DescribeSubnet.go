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
    cps "github.com/jdcloud-api/jdcloud-sdk-go/services/cps/models"
)

type DescribeSubnetRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的区域  */
    RegionId string `json:"regionId"`

    /* 可用区, 如cn-east-1a；可调用接口（describeRegiones）获取云物理服务器支持的可用区  */
    Az string `json:"az"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的区域 (Required)
 * param az: 可用区, 如cn-east-1a；可调用接口（describeRegiones）获取云物理服务器支持的可用区 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSubnetRequest(
    regionId string,
    az string,
) *DescribeSubnetRequest {

	return &DescribeSubnetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/subnet",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Az: az,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的区域 (Required)
 * param az: 可用区, 如cn-east-1a；可调用接口（describeRegiones）获取云物理服务器支持的可用区 (Required)
 */
func NewDescribeSubnetRequestWithAllParams(
    regionId string,
    az string,
) *DescribeSubnetRequest {

    return &DescribeSubnetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subnet",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Az: az,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSubnetRequestWithoutParam() *DescribeSubnetRequest {

    return &DescribeSubnetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subnet",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的区域(Required) */
func (r *DescribeSubnetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param az: 可用区, 如cn-east-1a；可调用接口（describeRegiones）获取云物理服务器支持的可用区(Required) */
func (r *DescribeSubnetRequest) SetAz(az string) {
    r.Az = az
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSubnetRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSubnetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSubnetResult `json:"result"`
}

type DescribeSubnetResult struct {
    Subnet cps.Subnet `json:"subnet"`
}