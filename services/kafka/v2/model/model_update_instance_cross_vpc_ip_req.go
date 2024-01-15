package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceCrossVpcIpReq struct {
	AdvertisedIpContents map[string]string `json:"advertised_ip_contents"`
}

func (o UpdateInstanceCrossVpcIpReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceCrossVpcIpReq struct{}"
	}

	return strings.Join([]string{"UpdateInstanceCrossVpcIpReq", string(data)}, " ")
}
