package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSlowlogDesensitizationSwitchRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`
}

func (o ShowSlowlogDesensitizationSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowlogDesensitizationSwitchRequest struct{}"
	}

	return strings.Join([]string{"ShowSlowlogDesensitizationSwitchRequest", string(data)}, " ")
}
