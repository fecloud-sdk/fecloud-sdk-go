package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateRouteTableRequest struct {
	RoutetableId string `json:"routetable_id"`

	Body *UpdateRoutetableReqBody `json:"body,omitempty"`
}

func (o UpdateRouteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRouteTableRequest struct{}"
	}

	return strings.Join([]string{"UpdateRouteTableRequest", string(data)}, " ")
}
