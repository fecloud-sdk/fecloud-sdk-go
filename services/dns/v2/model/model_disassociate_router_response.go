package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DisassociateRouterResponse struct {
	RouterId *string `json:"router_id,omitempty"`

	RouterRegion *string `json:"router_region,omitempty"`

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisassociateRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateRouterResponse struct{}"
	}

	return strings.Join([]string{"DisassociateRouterResponse", string(data)}, " ")
}
