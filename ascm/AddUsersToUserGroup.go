package ascm

import (
	"fmt"

	"github.com/devforma/apsara/core"
	"github.com/devforma/apsara/util"
)

type AddUsersToUserGroupRequest struct {
	core.EmbededRequest
}

type AddUsersToUserGroupResponse struct {
	core.EmbededResponse
}

func NewAddUsersToUserGroupRequest(userGroupId int64, loginName string) *AddUsersToUserGroupRequest {
	return &AddUsersToUserGroupRequest{
		core.EmbededRequest{
			Product: "ascm",
			Version: "2019-05-10",
			Action:  "AddUsersToUserGroup",
			Style:   core.RequestStyleRPC,
			BizQueries: map[string]string{
				"LoginNameList": fmt.Sprintf(`["%s"]`, loginName),
				"UserGroupId":   util.Int64ToString(userGroupId),
			},
		},
	}
}
