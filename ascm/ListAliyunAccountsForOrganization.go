package ascm

import (
	"github.com/devforma/apsara/core"
	"github.com/devforma/apsara/util"
	"github.com/tidwall/gjson"
)

type ListAliyunAccountsForOrganizationRequest struct {
	core.EmbededRequest
}

type ListAliyunAccountsForOrganizationResponse struct {
	core.EmbededResponse
}

func NewListAliyunAccountsForOrganizationRequest(regionId string, organizationId int64) *ListAliyunAccountsForOrganizationRequest {
	return &ListAliyunAccountsForOrganizationRequest{
		core.EmbededRequest{
			Product: "ascm",
			Version: "2019-05-10",
			Action:  "ListAliyunAccountsForOrganization",
			Style:   core.RequestStyleRPC,
			BizQueries: map[string]string{
				"RegionId":       regionId,
				"OrganizationId": util.Int64ToString(organizationId),
			},
		},
	}
}

type OrganizationAccount struct {
	AccessKeyId     string
	AccessKeySecret string
	AliyunId        string // odps project添加账号会用到
	PrimaryKey      string // oss uid使用
}

// 只有一级组织才有这个信息，下级组织调用该接口得到的都是其一级组织的信息
func (r *ListAliyunAccountsForOrganizationResponse) GetInfo() *OrganizationAccount {
	var account *OrganizationAccount
	r.Body.Get("data").ForEach(func(_, value gjson.Result) bool {
		account = &OrganizationAccount{
			AccessKeyId:     value.Get("accessKeyId").String(),
			AccessKeySecret: value.Get("accessKeySecret").String(),
			AliyunId:        value.Get("aliyunid").String(),
			PrimaryKey:      value.Get("primaryKey").String(),
		}

		return false
	})

	return account
}
