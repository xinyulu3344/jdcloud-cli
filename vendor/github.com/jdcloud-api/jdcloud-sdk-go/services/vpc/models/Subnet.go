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

package models


type Subnet struct {

    /* Subnet的Id (Optional) */
    SubnetId string `json:"subnetId"`

    /* 子网名称 (Optional) */
    SubnetName string `json:"subnetName"`

    /* 子网所属VPC的Id (Optional) */
    VpcId string `json:"vpcId"`

    /* 子网网段，vpc内子网网段不能重叠，cidr的取值范围：10.0.0.0/8、172.16.0.0/12和192.168.0.0/16及它们包含的子网，且子网掩码长度为16-28之间，如果VPC含有Cidr，则必须为VPC所在Cidr的子网 (Optional) */
    AddressPrefix string `json:"addressPrefix"`

    /* 子网可用ip数量 (Optional) */
    AvailableIpCount int `json:"availableIpCount"`

    /* 子网描述信息 (Optional) */
    Description string `json:"description"`

    /* 子网关联的路由表Id (Optional) */
    RouteTableId string `json:"routeTableId"`

    /* 子网关联的acl Id (Optional) */
    AclId string `json:"aclId"`

    /* 子网的起始地址，子网第1个地位为路由器网关保留，第2个地址为dhcp服务保留 (Optional) */
    StartIp string `json:"startIp"`

    /* 子网的结束地址，子网第1个地位为路由器网关保留，第2个地址为dhcp服务保留 (Optional) */
    EndIp string `json:"endIp"`

    /* 子网创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`
}
