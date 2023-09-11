package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DisassociateRouteTableRequest struct {
	RoutetableId string `json:"routetable_id"`

	Body *RoutetableAssociateReqbody `json:"body,omitempty"`
}

func (o DisassociateRouteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateRouteTableRequest struct{}"
	}

	return strings.Join([]string{"DisassociateRouteTableRequest", string(data)}, " ")
}
