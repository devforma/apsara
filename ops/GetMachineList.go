package ops

import (
	"strconv"

	"github.com/devforma/apsara/core"
	"github.com/tidwall/gjson"
)

type GetMachineListRequest struct {
	core.EmbededRequest
}

func NewGetMachineListRequest(pageNum, pageSize int) *GetMachineListRequest {
	return &GetMachineListRequest{
		core.EmbededRequest{
			Product:  "ASO",
			Version:  "2019-11-13",
			Action:   "getMachineList",
			Style:    core.RequestStyleRPC,
			Pathname: "/aso/v3/physicalInfo/getMachineList",
			BizQueries: map[string]string{
				"PageSize": strconv.Itoa(pageSize),
				"PageNum":  strconv.Itoa(pageNum),
				"IsFormat": "false",
			},
		},
	}
}

func (req *GetMachineListRequest) GetHeaders() map[string]string {
	headers := req.EmbededRequest.GetHeaders()
	headers["x-ascm-pass-through-mode"] = "true"
	return headers
}

type GetMachineListResponse struct {
	core.EmbededResponse
}

type GetMachineListResponseItem struct {
	Project    string
	Cluster    string
	OS         string
	SN         string
	Hostname   string
	IP         string
	HWCPU      int64
	HWMEM      int64
	HWHarddisk int64
	State      string
}

func (resp *GetMachineListResponse) GetMachineList() []GetMachineListResponseItem {
	var items []GetMachineListResponseItem

	dataStr := resp.Body.Get("data").String()
	if dataStr == "" {
		return items
	}

	gjson.Parse(dataStr).ForEach(func(_, value gjson.Result) bool {
		if !value.Get("is_vm").Bool() {
			item := GetMachineListResponseItem{
				Project:    value.Get("project").String(),
				Cluster:    value.Get("cluster").String(),
				OS:         value.Get("os").String(),
				SN:         value.Get("sn").String(),
				Hostname:   value.Get("hostname").String(),
				IP:         value.Get("ip").String(),
				HWCPU:      value.Get("hw_cpu").Int(),
				HWMEM:      value.Get("hw_mem").Int(),
				HWHarddisk: value.Get("hw_harddisk").Int(),
				State:      value.Get("state").String(),
			}

			items = append(items, item)
		}

		return true
	})

	return items
}
