package odps

import (
	"fmt"

	"github.com/devforma/apsara/core"
	"github.com/devforma/apsara/util"
)

type CreateOdpsEngineRequest struct {
	core.EmbededRequest
}

type CreateOdpsEngineResponse struct {
	core.EmbededResponse
}

// NewCreateOdpsEngineRequest 调用前先判断是否存在重名engine
func NewCreateOdpsEngineRequest(organizationId, resourceGroupId int64, regionId, vpcId, cluster, taskUserPrimaryKey, name, desc string, cuId, storageGB int64) *CreateOdpsEngineRequest {
	detail := `[{"cluster":"%s","core_arch":"x86_64","cu":"%d","project":"odps","disk":%f,"isDefault":1}]`

	return &CreateOdpsEngineRequest{
		core.EmbededRequest{
			Product: "ascm",
			Version: "2019-05-10",
			Action:  "CreateOdpsEngine",
			Style:   core.RequestStyleRPC,
			BizQueries: map[string]string{
				"ClusterDetail":   fmt.Sprintf(detail, cluster, cuId, float64(storageGB)/1024),
				"OdpsName":        name,
				"Description":     desc,
				"resourceGroupId": util.Int64ToString(resourceGroupId),
				"OrganizationId":  util.Int64ToString(organizationId),
				"VpcTunnelIdList": regionId + "_" + vpcId,
				"TaskPk":          taskUserPrimaryKey,
				"AsArchitechture": "x86",
				"_noSign":         "true",
			},
		},
	}
}
