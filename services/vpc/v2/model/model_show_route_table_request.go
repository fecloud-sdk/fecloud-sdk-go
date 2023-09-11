package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowRouteTableRequest struct {
	RoutetableId string `json:"routetable_id"`
}

func (o ShowRouteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRouteTableRequest struct{}"
	}

	return strings.Join([]string{"ShowRouteTableRequest", string(data)}, " ")
}
