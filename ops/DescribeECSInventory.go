package ops

import (
	"github.com/devforma/apsara/core"
	"github.com/tidwall/gjson"
)

type DescribeECSInventoryRequest struct {
	core.EmbededRequest
}

type DescribeECSInventoryResponse struct {
	core.EmbededResponse
}

func NewDescribeECSInventoryRequest() *DescribeECSInventoryRequest {
	return &DescribeECSInventoryRequest{
		core.EmbededRequest{
			Product: "opsAPI",
			Version: "2018-01-22",
			Action:  "DescribeECSInventory",
			Style:   core.RequestStyleRPC,
		},
	}
}

type DescribeECSInventoryResponseItem struct {
	ZoneId            string
	AllTotalCPU       int64
	AllSoldCPU        int64
	AllVendibleCPU    int64
	AllTotalMemory    int64
	AllSoldMemory     int64
	AllVendibleMemory int64
	AllTotalGPU       int64
	AllSoldGPU        int64
	AllVendibleGPU    int64
}

func (resp *DescribeECSInventoryResponse) GetInventory() []DescribeECSInventoryResponseItem {
	var items []DescribeECSInventoryResponseItem
	resp.Body.Get("zones").ForEach(func(_, value gjson.Result) bool {
		item := DescribeECSInventoryResponseItem{
			ZoneId:            value.Get("zoneId").String(),
			AllTotalCPU:       value.Get("allTotalCPU").Int(),
			AllSoldCPU:        value.Get("allSoldCPU").Int(),
			AllVendibleCPU:    value.Get("allVendibleCPU").Int(),
			AllTotalMemory:    value.Get("allTotalMemory").Int(),
			AllSoldMemory:     value.Get("allSoldMemory").Int(),
			AllVendibleMemory: value.Get("allVendibleMemory").Int(),
			AllTotalGPU:       value.Get("allTotalGPU").Int(),
			AllSoldGPU:        value.Get("allSoldGPU").Int(),
			AllVendibleGPU:    value.Get("allVendibleGPU").Int(),
		}

		items = append(items, item)
		return true
	})

	return items
}
