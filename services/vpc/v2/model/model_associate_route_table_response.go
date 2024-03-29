package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AssociateRouteTableResponse struct {
	Routetable     *RouteTableResp `json:"routetable,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o AssociateRouteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRouteTableResponse struct{}"
	}

	return strings.Join([]string{"AssociateRouteTableResponse", string(data)}, " ")
}
