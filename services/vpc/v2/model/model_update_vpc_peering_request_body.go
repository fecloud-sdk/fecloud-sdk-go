package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateVpcPeeringRequestBody struct {
	Peering *UpdateVpcPeeringOption `json:"peering"`
}

func (o UpdateVpcPeeringRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcPeeringRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVpcPeeringRequestBody", string(data)}, " ")
}
