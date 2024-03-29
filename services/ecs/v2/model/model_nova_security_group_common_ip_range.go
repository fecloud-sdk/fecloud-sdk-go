package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NovaSecurityGroupCommonIpRange struct {
	Cidr *string `json:"cidr,omitempty"`
}

func (o NovaSecurityGroupCommonIpRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaSecurityGroupCommonIpRange struct{}"
	}

	return strings.Join([]string{"NovaSecurityGroupCommonIpRange", string(data)}, " ")
}
