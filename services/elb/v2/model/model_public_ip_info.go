package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PublicIpInfo struct {
	PublicipId string `json:"publicip_id"`

	PublicipAddress string `json:"publicip_address"`

	IpVersion int32 `json:"ip_version"`
}

func (o PublicIpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIpInfo struct{}"
	}

	return strings.Join([]string{"PublicIpInfo", string(data)}, " ")
}
