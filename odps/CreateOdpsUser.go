package odps

import (
	"github.com/devforma/apsara/core"
	"github.com/devforma/apsara/util"
)

type CreateOdpsUserRequest struct {
	core.EmbededRequest
}

type CreateOdpsUserResponse struct {
	core.EmbededResponse
}

// NewCreateOdpsUserRequest 创建的云账号和ascm AddUser不是同一类账号，互相没有关系
func NewCreateOdpsUserRequest(organizationId int64, userName string) *CreateOdpsUserRequest {
	return &CreateOdpsUserRequest{
		core.EmbededRequest{
			Product: "ascm",
			Version: "2019-05-10",
			Action:  "CreateOdpsUser",
			Style:   core.RequestStyleRPC,
			BizQueries: map[string]string{
				"OrganizationId": util.Int64ToString(organizationId),
				"UserName":       userName,
			},
		},
	}
}
