package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RejectVpcPeeringRequest struct {
	PeeringId string `json:"peering_id"`
}

func (o RejectVpcPeeringRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RejectVpcPeeringRequest struct{}"
	}

	return strings.Join([]string{"RejectVpcPeeringRequest", string(data)}, " ")
}
