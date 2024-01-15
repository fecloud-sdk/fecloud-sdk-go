package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AssociateRouterResponse struct {
	RouterId *string `json:"router_id,omitempty"`

	RouterRegion *string `json:"router_region,omitempty"`

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRouterResponse struct{}"
	}

	return strings.Join([]string{"AssociateRouterResponse", string(data)}, " ")
}
