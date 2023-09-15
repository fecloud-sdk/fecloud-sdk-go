package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowIpWhitelistRequest Request Object
type ShowIpWhitelistRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowIpWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpWhitelistRequest struct{}"
	}

	return strings.Join([]string{"ShowIpWhitelistRequest", string(data)}, " ")
}
