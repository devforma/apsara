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
// 其中taskPK为任务账号，应该是用于dataworks里面使用，目前拿不到aksk，不适用于odpscmd和odps sdk中aksk方式
// 向odps添加ascm云账号时，需要使用odps sdk或者odpscmd来执行 add user RAM$组织云账号id:组织下用户loginName; 例如：add user RAM$ascm-org-1602732819130:citybrain_testuser3; 然后进行grant权限配置
// acl权限控制最佳实践：CREATE ROLE worker; GRANT 权限 TO ROLE worker; GRANT worker TO RAM$ascm-org-1602732819130:citybrain_testuser3; 文档见https://help.aliyun.com/apsara/enterprise/v_3_16_2_20220708/odps/enterprise-ascm-user-guide/acl-authorization-actions.html
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
				"TaskPk":          taskUserPrimaryKey,
				// "VpcTunnelIdList": regionId + "_" + vpcId,
				// "AsArchitechture": "x86",
				// "_noSign":         "true",
			},
		},
	}
}
