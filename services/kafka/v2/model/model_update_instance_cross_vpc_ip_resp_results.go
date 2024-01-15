package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceCrossVpcIpRespResults struct {
	AdvertisedIp *string `json:"advertised_ip,omitempty"`

	Success *bool `json:"success,omitempty"`

	Ip *string `json:"ip,omitempty"`
}

func (o UpdateInstanceCrossVpcIpRespResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceCrossVpcIpRespResults struct{}"
	}

	return strings.Join([]string{"UpdateInstanceCrossVpcIpRespResults", string(data)}, " ")
}
