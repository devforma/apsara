package odps

import (
	"github.com/devforma/apsara/core"
	"github.com/devforma/apsara/util"
	"github.com/tidwall/gjson"
)

type ListOdpsCUsRequest struct {
	core.EmbededRequest
}

type ListOdpsCUsResponse struct {
	core.EmbededResponse
}

func NewListOdpsCUsRequest(organizationId, resourceGroupId int64, cluster string) *ListOdpsCUsRequest {
	return &ListOdpsCUsRequest{
		core.EmbededRequest{
			Product:        "ascm",
			Version:        "2019-05-10",
			Action:         "ListOdpsCus",
			Style:          core.RequestStyleRPC,
			OrganizationID: util.Int64ToString(organizationId),
			BizQueries: map[string]string{
				"ClusterName":   cluster,
				"PageNum":       "1",
				"PageSize":      "999",
				"Project":       "odps",
				"Department":    util.Int64ToString(organizationId),
				"ResourceGroup": util.Int64ToString(resourceGroupId),
			},
		},
	}
}

type CU struct {
	Id            int64
	QuotaGroup    string
	CpuMemoryRate string
	MinCU         int64
	MaxCU         int64
	MinMem        int64
	MaxMem        int64
	Mem           int64
	MinCpu        int64
	MaxCpu        int64
	CPU           int64
}

func (resp *ListOdpsCUsResponse) GetCUs() []CU {
	var cus []CU
	resp.Body.Get("data").ForEach(func(_, value gjson.Result) bool {
		cu := CU{
			Id:            value.Get("id").Int(),
			QuotaGroup:    value.Get("quotagroup").String(),
			CpuMemoryRate: value.Get("cpu_memory_rate").String(),
			MinCU:         value.Get("min_cu").Int(),
			MaxCU:         value.Get("max_cu").Int(),
			MinMem:        value.Get("min_mem").Int(),
			MaxMem:        value.Get("max_mem").Int(),
			Mem:           value.Get("mem").Int(),
			MinCpu:        value.Get("min_cpu").Int(),
			MaxCpu:        value.Get("max_cpu").Int(),
			CPU:           value.Get("cpu").Int(),
		}
		cus = append(cus, cu)
		return true
	})

	return cus
}
