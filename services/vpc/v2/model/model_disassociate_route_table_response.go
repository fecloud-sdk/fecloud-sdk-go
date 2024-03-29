package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DisassociateRouteTableResponse struct {
	Routetable     *RouteTableResp `json:"routetable,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o DisassociateRouteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateRouteTableResponse struct{}"
	}

	return strings.Join([]string{"DisassociateRouteTableResponse", string(data)}, " ")
}
