package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AssociateRouterRequest struct {
	ZoneId string `json:"zone_id"`

	Body *AssociateRouterRequestBody `json:"body,omitempty"`
}

func (o AssociateRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRouterRequest struct{}"
	}

	return strings.Join([]string{"AssociateRouterRequest", string(data)}, " ")
}
