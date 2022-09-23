package odps

import (
	"github.com/devforma/apsara/core"
	"github.com/devforma/apsara/util"
	"github.com/tidwall/gjson"
)

type GetOdpsUserInfoRequest struct {
	core.EmbededRequest
}

type GetOdpsUserInfoResponse struct {
	core.EmbededResponse
}

func NewGetOdpsUserInfoRequest(organizationId int64, userName string) *GetOdpsUserInfoRequest {
	return &GetOdpsUserInfoRequest{
		core.EmbededRequest{
			Product:        "ascm",
			Version:        "2019-05-10",
			Action:         "GetOdpsUserList",
			Style:          core.RequestStyleRPC,
			OrganizationID: util.Int64ToString(organizationId),
			BizQueries: map[string]string{
				"UserName": userName,
			},
			CachedParams: map[string]string{
				"UserName": userName,
			},
		},
	}
}

type OdpsUserInfo struct {
	UserName   string
	PrimaryKey string
	UserId     string
}

func (resp *GetOdpsUserInfoResponse) GetOdpsUserInfo() *OdpsUserInfo {
	var found *OdpsUserInfo
	resp.Body.Get("data").ForEach(func(_, value gjson.Result) bool {
		if userName := value.Get("userName").String(); userName == resp.CachedRequestParams["UserName"] {
			found = &OdpsUserInfo{
				UserName:   userName,
				PrimaryKey: value.Get("aasPk").String(),
				UserId:     value.Get("userId").String(),
			}
			return false
		}
		return true
	})

	return found
}
