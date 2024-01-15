package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePrivateSnatOption struct {
	GatewayId string `json:"gateway_id"`

	Cidr *string `json:"cidr,omitempty"`

	VirsubnetId *string `json:"virsubnet_id,omitempty"`

	Description *string `json:"description,omitempty"`

	TransitIpIds []string `json:"transit_ip_ids"`
}

func (o CreatePrivateSnatOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateSnatOption struct{}"
	}

	return strings.Join([]string{"CreatePrivateSnatOption", string(data)}, " ")
}
