package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteVpcRouteRequest struct {
	RouteId string `json:"route_id"`
}

func (o DeleteVpcRouteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcRouteRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpcRouteRequest", string(data)}, " ")
}
