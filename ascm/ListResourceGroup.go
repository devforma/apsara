package ascm

import (
	"github.com/devforma/apsara/core"
	"github.com/devforma/apsara/util"
)

type ListResourceGroupRequest struct {
	core.EmbededRequest
}

type ListResourceGroupResponse struct {
	core.EmbededResponse
}

func NewListResourceGroupRequest(organizationId int64) *ListResourceGroupRequest {
	return &ListResourceGroupRequest{
		core.EmbededRequest{
			Product: "ascm",
			Version: "2019-05-10",
			Action:  "ListResourceGroup",
			Style:   core.RequestStyleRPC,
			BizQueries: map[string]string{
				"OrganizationId": util.Int64ToString(organizationId),
				"pageNumber":     "1",
				"pageSize":       "10",
			},
		},
	}
}

func (r *ListResourceGroupResponse) GetID() int64 {
	return r.Body.Get("data.0.id").Int()
}
