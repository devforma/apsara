package odps

import (
	"github.com/devforma/apsara/core"
	"github.com/devforma/apsara/util"
)

type CreateUpdateOdpsCURequest struct {
	core.EmbededRequest
}

type CreateUpdateOdpsCUResponse struct {
	core.EmbededResponse
}

func NewCreateUpdateOdpsCURequest(organizationId, resourceGroupId int64, regionId, cluster, name string, cuNum int64) *CreateUpdateOdpsCURequest {
	return &CreateUpdateOdpsCURequest{
		core.EmbededRequest{
			Product: "ascm",
			Version: "2019-05-10",
			Action:  "CreateUpdateOdpsCu",
			Style:   core.RequestStyleRPC,
			BizQueries: map[string]string{
				"OrganizationId":  util.Int64ToString(organizationId),
				"ResourceGroupId": util.Int64ToString(resourceGroupId),
				"RegionId":        regionId,
				"RegionName":      regionId,
				"CuName":          name,
				"ClusterName":     cluster,
				"Share":           "2", // 0同资源集共享 1本组织及下级组织共享 2同组织共享
				"CuNum":           util.Int64ToString(cuNum),
			},
		},
	}
}
