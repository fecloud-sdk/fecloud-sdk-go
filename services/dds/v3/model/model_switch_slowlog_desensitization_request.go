package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SwitchSlowlogDesensitizationRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Status string `json:"status"`
}

func (o SwitchSlowlogDesensitizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSlowlogDesensitizationRequest struct{}"
	}

	return strings.Join([]string{"SwitchSlowlogDesensitizationRequest", string(data)}, " ")
}
