package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowIpWhitelistRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowIpWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpWhitelistRequest struct{}"
	}

	return strings.Join([]string{"ShowIpWhitelistRequest", string(data)}, " ")
}
