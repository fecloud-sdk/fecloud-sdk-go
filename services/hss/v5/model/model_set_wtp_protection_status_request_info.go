package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetWtpProtectionStatusRequestInfo struct {
	Status *bool `json:"status,omitempty"`

	HostIdList *[]string `json:"host_id_list,omitempty"`

	ResourceId *string `json:"resource_id,omitempty"`

	ChargingMode *string `json:"charging_mode,omitempty"`
}

func (o SetWtpProtectionStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetWtpProtectionStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"SetWtpProtectionStatusRequestInfo", string(data)}, " ")
}
