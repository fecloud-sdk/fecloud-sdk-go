package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DisassociateRouterRequest struct {
	ZoneId string `json:"zone_id"`

	Body *DisassociaterouterRequestBody `json:"body,omitempty"`
}

func (o DisassociateRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateRouterRequest struct{}"
	}

	return strings.Join([]string{"DisassociateRouterRequest", string(data)}, " ")
}
