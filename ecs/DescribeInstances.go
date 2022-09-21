package ecs

import (
	"time"

	"github.com/devforma/apsara/core"
	"github.com/devforma/apsara/util"
	"github.com/tidwall/gjson"
)

type DescribeInstancesRequest struct {
	core.EmbededRequest
}

func NewDescribeInstancesRequest(instanceIds []string) *DescribeInstancesRequest {
	return &DescribeInstancesRequest{
		core.EmbededRequest{
			Product: "Ecs",
			Version: "2014-05-26",
			Action:  "DescribeInstances",
			Style:   core.RequestStyleRPC,
			BizQueries: map[string]string{
				"PageSize":    "100",
				"InstanceIds": util.JsonMarshal(instanceIds),
			},
		},
	}
}

type DescribeInstancesResponse struct {
	core.EmbededResponse
}

type DescribeInstancesResponseItem struct {
	Memory             int64
	InstanceChargeType string
	Cpu                int64
	OSNameEn           string
	InstanceBandwidth  int64
	InstanceType       string
	Department         int64
	ResourceGroup      int64
	CreationTime       time.Time
	StartTime          time.Time
	ExpiredTime        time.Time
	ImageId            string
	Status             string
	InstanceId         string
	InstanceName       string
	SecurityGroupIds   []string
	IpAddress          string
	VpcId              string
	VSwitchId          string
}

func (resp *DescribeInstancesResponse) GetInstances() []DescribeInstancesResponseItem {
	var items []DescribeInstancesResponseItem

	gjson.GetBytes(resp.Body, "Instances.Instance").ForEach(func(_, value gjson.Result) bool {
		item := DescribeInstancesResponseItem{
			Memory:             value.Get("Memory").Int(),
			InstanceChargeType: value.Get("InstanceChargeType").String(),
			Cpu:                value.Get("Cpu").Int(),
			OSNameEn:           value.Get("OSNameEn").String(),
			InstanceBandwidth:  value.Get("InstanceBandwidthRx").Int(),
			InstanceType:       value.Get("InstanceType").String(),
			Department:         value.Get("Department").Int(),
			ResourceGroup:      value.Get("ResourceGroup").Int(),
			CreationTime:       parseTime(value.Get("CreationTime").String()),
			StartTime:          parseTime(value.Get("StartTime").String()),
			ExpiredTime:        parseTime(value.Get("ExpiredTime").String()),
			ImageId:            value.Get("ImageId").String(),
			Status:             value.Get("Status").String(),
			InstanceId:         value.Get("InstanceId").String(),
			InstanceName:       value.Get("InstanceName").String(),
			IpAddress:          value.Get("VpcAttributes.PrivateIpAddress.IpAddress.0").String(),
			VpcId:              value.Get("VpcAttributes.VpcId").String(),
			VSwitchId:          value.Get("VpcAttributes.VSwitchId").String(),
		}
		items = append(items, item)

		return true
	})

	return items
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04Z", timeStr)
	return t
}
