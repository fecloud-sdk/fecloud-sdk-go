package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetRaspSwitchRequestInfo struct {
	HostIdList *[]string `json:"host_id_list,omitempty"`

	Status *bool `json:"status,omitempty"`
}

func (o SetRaspSwitchRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRaspSwitchRequestInfo struct{}"
	}

	return strings.Join([]string{"SetRaspSwitchRequestInfo", string(data)}, " ")
}
