package ascm

import (
	"github.com/devforma/apsara/core"
	"github.com/devforma/apsara/util"
)

type AddUserRequest struct {
	core.EmbededRequest
}

type AddUserResponse struct {
	core.EmbededResponse
}

// NewAddUserRequest 创建用户，默认登录策略为允许登录ascm后台，但初始密码需要在ascm后台查看用户信息处获取，用初始密码后登录需重置密码
func NewAddUserRequest(organizationId int64, loginName, displayName, phone, email string) *AddUserRequest {
	return &AddUserRequest{
		core.EmbededRequest{
			Product: "ascm",
			Version: "2019-05-10",
			Action:  "AddUser",
			Style:   core.RequestStyleRPC,
			BizQueries: map[string]string{
				"CellphoneNum":     phone,
				"DisplayName":      displayName,
				"Email":            email,
				"LoginName":        loginName,
				"MobileNationCode": "86",
				"OrganizationId":   util.Int64ToString(organizationId),
			},
		},
	}
}

type UserInfo struct {
	ID              int64
	PrimaryKey      string
	AccessKeyID     string
	AccessKeySecret string
}

func (resp *AddUserResponse) GetUserInfo() *UserInfo {
	primaryKey := resp.Body.Get("data.primaryKey").String()
	if primaryKey == "" {
		return nil
	}

	return &UserInfo{
		ID:              resp.Body.Get("data.id").Int(),
		PrimaryKey:      primaryKey,
		AccessKeyID:     resp.Body.Get("data.accessKeys.0.accesskeyId").String(),
		AccessKeySecret: resp.Body.Get("data.accessKeys.0.accesskeySecret").String(),
	}
}
