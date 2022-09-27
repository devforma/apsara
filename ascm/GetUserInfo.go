package ascm

import "github.com/devforma/apsara/core"

type GetUserInfoRequest struct {
	core.EmbededRequest
}

type GetUserInfoResponse struct {
	core.EmbededResponse
}

// NewGetUserInfoRequest 获取用户信息，只有client的aksk和loginName是用一个用户，则返回结果中有aksk信息，否则没有
func NewGetUserInfoRequest(loginName string) *GetUserInfoRequest {
	return &GetUserInfoRequest{
		core.EmbededRequest{
			Product: "ascm",
			Version: "2019-05-10",
			Action:  "GetUserInfo",
			Style:   core.RequestStyleRPC,
			BizQueries: map[string]string{
				"loginName": loginName,
			},
		},
	}
}
