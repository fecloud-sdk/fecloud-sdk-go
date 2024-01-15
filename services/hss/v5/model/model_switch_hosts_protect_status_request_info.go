package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SwitchHostsProtectStatusRequestInfo struct {
	Version string `json:"version"`

	ChargingMode *string `json:"charging_mode,omitempty"`

	ResourceId *string `json:"resource_id,omitempty"`

	HostIdList []string `json:"host_id_list"`

	Tags *[]TagInfo `json:"tags,omitempty"`
}

func (o SwitchHostsProtectStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchHostsProtectStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"SwitchHostsProtectStatusRequestInfo", string(data)}, " ")
}
