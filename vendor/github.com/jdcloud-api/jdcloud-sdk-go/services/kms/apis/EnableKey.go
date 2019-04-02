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

type EnableKeyRequest struct {

    core.JDCloudRequest

    /* 密钥ID  */
    KeyId string `json:"keyId"`
}

/*
 * param keyId: 密钥ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEnableKeyRequest(
    keyId string,
) *EnableKeyRequest {

	return &EnableKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/key/{keyId}:enable",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        KeyId: keyId,
	}
}

/*
 * param keyId: 密钥ID (Required)
 */
func NewEnableKeyRequestWithAllParams(
    keyId string,
) *EnableKeyRequest {

    return &EnableKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/key/{keyId}:enable",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        KeyId: keyId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEnableKeyRequestWithoutParam() *EnableKeyRequest {

    return &EnableKeyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/key/{keyId}:enable",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param keyId: 密钥ID(Required) */
func (r *EnableKeyRequest) SetKeyId(keyId string) {
    r.KeyId = keyId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EnableKeyRequest) GetRegionId() string {
    return ""
}

type EnableKeyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EnableKeyResult `json:"result"`
}

type EnableKeyResult struct {
}