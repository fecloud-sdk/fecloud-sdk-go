package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowIpGroupRequest struct {
	IpgroupId string `json:"ipgroup_id"`
}

func (o ShowIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowIpGroupRequest", string(data)}, " ")
}
