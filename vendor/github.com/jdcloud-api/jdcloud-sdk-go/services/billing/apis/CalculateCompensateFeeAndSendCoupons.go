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
    billing "github.com/jdcloud-api/jdcloud-sdk-go/services/billing/models"
)

type CalculateCompensateFeeAndSendCouponsRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 用户pin  */
    Pin string `json:"pin"`

    /* 资源ID  */
    ResourceIds []string `json:"resourceIds"`
}

/*
 * param regionId:  (Required)
 * param pin: 用户pin (Required)
 * param resourceIds: 资源ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCalculateCompensateFeeAndSendCouponsRequest(
    regionId string,
    pin string,
    resourceIds []string,
) *CalculateCompensateFeeAndSendCouponsRequest {

	return &CalculateCompensateFeeAndSendCouponsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/price/calculateCompensateFeeAndSendCoupons",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Pin: pin,
        ResourceIds: resourceIds,
	}
}

/*
 * param regionId:  (Required)
 * param pin: 用户pin (Required)
 * param resourceIds: 资源ID (Required)
 */
func NewCalculateCompensateFeeAndSendCouponsRequestWithAllParams(
    regionId string,
    pin string,
    resourceIds []string,
) *CalculateCompensateFeeAndSendCouponsRequest {

    return &CalculateCompensateFeeAndSendCouponsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/price/calculateCompensateFeeAndSendCoupons",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
        ResourceIds: resourceIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCalculateCompensateFeeAndSendCouponsRequestWithoutParam() *CalculateCompensateFeeAndSendCouponsRequest {

    return &CalculateCompensateFeeAndSendCouponsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/price/calculateCompensateFeeAndSendCoupons",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *CalculateCompensateFeeAndSendCouponsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: 用户pin(Required) */
func (r *CalculateCompensateFeeAndSendCouponsRequest) SetPin(pin string) {
    r.Pin = pin
}

/* param resourceIds: 资源ID(Required) */
func (r *CalculateCompensateFeeAndSendCouponsRequest) SetResourceIds(resourceIds []string) {
    r.ResourceIds = resourceIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CalculateCompensateFeeAndSendCouponsRequest) GetRegionId() string {
    return r.RegionId
}

type CalculateCompensateFeeAndSendCouponsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CalculateCompensateFeeAndSendCouponsResult `json:"result"`
}

type CalculateCompensateFeeAndSendCouponsResult struct {
    Pin string `json:"pin"`
    TotalCompensateFee int `json:"totalCompensateFee"`
    ResourceList []billing.OrderCompensateFeeVo `json:"resourceList"`
}