package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EniNetwork struct {
	EniSubnetId string `json:"eniSubnetId"`

	EniSubnetCIDR string `json:"eniSubnetCIDR"`

	Subnets []NetworkSubnet `json:"subnets"`
}

func (o EniNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EniNetwork struct{}"
	}

	return strings.Join([]string{"EniNetwork", string(data)}, " ")
}
