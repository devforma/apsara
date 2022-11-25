package ack

import (
	"github.com/devforma/apsara/core"
	"github.com/devforma/apsara/util"
)

type DescribeClustersRequest struct {
	core.EmbededRequest
}

type DescribeClustersResponse struct {
	core.EmbededResponse
}

func NewDescribeClustersRequest(regionId string, organizationId int64, clusterId string) *DescribeClustersRequest {
	return &DescribeClustersRequest{
		core.EmbededRequest{
			Product:        "CS",
			Version:        "2015-12-15",
			Action:         "DescribeClusters",
			Style:          core.RequestStyleRPC,
			OrganizationID: util.Int64ToString(organizationId),
			BizQueries: map[string]string{
				"RegionName": regionId,
				"RegionId":   regionId,
			},
		},
	}
}
