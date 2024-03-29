package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteVpcPeeringRequest struct {
	PeeringId string `json:"peering_id"`
}

func (o DeleteVpcPeeringRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcPeeringRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpcPeeringRequest", string(data)}, " ")
}
