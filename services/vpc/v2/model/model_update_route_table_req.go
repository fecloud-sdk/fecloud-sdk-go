package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateRouteTableReq struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Routes map[string][]RouteTableRoute `json:"routes,omitempty"`
}

func (o UpdateRouteTableReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRouteTableReq struct{}"
	}

	return strings.Join([]string{"UpdateRouteTableReq", string(data)}, " ")
}
