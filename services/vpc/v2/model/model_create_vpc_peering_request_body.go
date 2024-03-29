package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateVpcPeeringRequestBody struct {
	Peering *CreateVpcPeeringOption `json:"peering"`
}

func (o CreateVpcPeeringRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcPeeringRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVpcPeeringRequestBody", string(data)}, " ")
}
