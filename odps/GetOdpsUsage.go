package odps

import (
	"encoding/json"
	"time"

	"github.com/devforma/apsara/core"
	"github.com/devforma/apsara/util"
)

type GetOdpsUsageRequest struct {
	core.EmbededRequest
}

type GetOdpsUsageResponse struct {
	core.EmbededResponse
}

func NewGetOdpsUsageRequest(cluster string, startTime, endTime time.Time) *GetOdpsUsageRequest {
	return &GetOdpsUsageRequest{
		core.EmbededRequest{
			Style:    core.RequestStyleBCC,
			Pathname: "/gateway/v2/apps/eodps/faas/faas_odps/service/fuxi/get_cluster_quota_history_info",
			BizQueries: map[string]string{
				"stime":   util.Int64ToString(startTime.Unix()),
				"etime":   util.Int64ToString(endTime.Unix()),
				"cluster": cluster,
				// "keys":    "cpu_used",
			},
		},
	}
}

type OdpsUsage struct {
	Time     time.Time
	CPUTotal int64
	CPUUsed  int64
	MemTotal int64
	MemUsed  int64
	GPUTotal int64
	GPUUsed  int64
}

type respData struct {
	Data struct {
		CPUTotal [][]int64 `json:"cpu_total"`
		CPUUsed  [][]int64 `json:"cpu_used"`
		MemTotal [][]int64 `json:"mem_total"`
		MemUsed  [][]int64 `json:"mem_used"`
		GPUTotal [][]int64 `json:"gpu_total"`
		GPUUsed  [][]int64 `json:"gpu_used"`
	} `json:"data"`
}

func (resp *GetOdpsUsageResponse) GetUsage() []OdpsUsage {
	var usageList []OdpsUsage

	var data respData
	if err := json.Unmarshal(util.StringToBytes(resp.Body.Raw), &data); err != nil {
		return usageList
	}

	for i := 0; i < len(data.Data.CPUTotal); i++ {
		usageList = append(usageList, OdpsUsage{
			Time:     time.Unix(data.Data.CPUTotal[i][0]/1000, 0),
			CPUTotal: data.Data.CPUTotal[i][1],
			CPUUsed:  data.Data.CPUUsed[i][1],
			MemTotal: data.Data.MemTotal[i][1],
			MemUsed:  data.Data.MemUsed[i][1],
			GPUTotal: data.Data.GPUTotal[i][1],
			GPUUsed:  data.Data.GPUUsed[i][1],
		})
	}

	return usageList
}
