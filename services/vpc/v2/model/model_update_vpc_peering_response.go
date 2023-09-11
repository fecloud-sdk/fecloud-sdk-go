package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateVpcPeeringResponse struct {
	Peering        *VpcPeering `json:"peering,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateVpcPeeringResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcPeeringResponse struct{}"
	}

	return strings.Join([]string{"UpdateVpcPeeringResponse", string(data)}, " ")
}
