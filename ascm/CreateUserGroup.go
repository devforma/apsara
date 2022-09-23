package ascm

import (
	"github.com/devforma/apsara/core"
	"github.com/devforma/apsara/util"
)

type CreateUserGroupRequest struct {
	core.EmbededRequest
}

type CreateUserGroupResponse struct {
	core.EmbededResponse
}

func NewCreateUserGroupRequest(organizationId int64, name, desc string) *CreateUserGroupRequest {
	return &CreateUserGroupRequest{
		core.EmbededRequest{
			Product: "ascm",
			Version: "2019-05-10",
			Action:  "CreateUserGroup",
			Style:   core.RequestStyleRPC,
			BizQueries: map[string]string{
				"GroupName":      name,
				"OrganizationId": util.Int64ToString(organizationId),
				"Description":    desc,
			},
		},
	}
}

func (resp *CreateUserGroupResponse) GetUserGroupId() int64 {
	return resp.Body.Get("data.id").Int()
}
