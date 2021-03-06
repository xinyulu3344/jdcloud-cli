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


type MetricDataCm struct {

    /* 命名空间 ，长度不超过255字节，只允许英文、数字、下划线_、点., [0-9][a-z] [A-Z] [. _ ]，  其它会返回err  */
    Namespace string `json:"namespace"`

    /* 监控指标名称，长度不超过255字节，只允许英文、数字、下划线_、点.,  [0-9][a-z] [A-Z] [. _ ]， 其它会返回err  */
    Metric string `json:"metric"`

    /* 数据维度，数据类型为map类型，支持最少一个，最多五个标签，总长度不大于255字节，只允许英文、数字、下划线_、点., [0-9][a-z] [A-Z] [. _ ]，  其它会返回err  */
    Dimensions interface{} `json:"dimensions"`

    /* 上报数据点的时间戳,只支持10位，秒级时间戳，不能写入过去30天的时间  */
    Timestamp int64 `json:"timestamp"`

    /* 数据上报类型，1为原始值，2为聚合数据。当上报聚合数据时，建议为60s的周期时行上报、否则无法正常查询  */
    Type int `json:"type"`

    /* 指标值集合，数据类型必须为map类型，key为数据类型，value为数据值，当type=1时，key只能为”value”，上报的是原始值，当type=2时，K的值可以为"avg","sum","last","max","min","count"，只支持以上类型，否则会报错，value内容为整型或浮点型数字，最大值为9223372036854775807，count只支持>=0的数  */
    Values interface{} `json:"values"`
}
