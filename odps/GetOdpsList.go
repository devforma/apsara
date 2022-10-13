package odps

import (
	"github.com/devforma/apsara/core"
	"github.com/devforma/apsara/util"
	"github.com/tidwall/gjson"
)

type GetOdpsEngineListRequest struct {
	core.EmbededRequest
}

type GetOdpsEngineListResponse struct {
	core.EmbededResponse
}

// NewGetOdpsEngineListRequest 获取odps projects列表
func NewGetOdpsEngineListRequest(organizationId, resourceGroupId int64, pageNum, pageSize int64) *GetOdpsEngineListRequest {
	return &GetOdpsEngineListRequest{
		core.EmbededRequest{
			Product:         "ascm",
			Version:         "2019-05-10",
			Action:          "GetOdpsEngineList",
			Style:           core.RequestStyleRPC,
			OrganizationID:  util.Int64ToString(organizationId),
			ResourceGroupID: util.Int64ToString(resourceGroupId),
			BizQueries: map[string]string{
				"CurrentPage": util.Int64ToString(pageNum),
				"PageSize":    util.Int64ToString(pageSize),
			},
		},
	}
}

type OdpsEngineListItem struct {
	Name        string
	Description string
	Id          int64
	EngineId    int64
	State       string
}

func (resp *GetOdpsEngineListResponse) GetOdpsEngineList() []OdpsEngineListItem {
	var items []OdpsEngineListItem
	resp.Body.Get("data").ForEach(func(_, value gjson.Result) bool {
		item := OdpsEngineListItem{
			Name:        value.Get("odpsName").String(),
			Description: value.Get("description").String(),
			EngineId:    value.Get("engineId").Int(),
			State:       value.Get("state").String(),
			Id:          value.Get("id").Int(),
		}
		items = append(items, item)

		return true
	})

	return items
}
