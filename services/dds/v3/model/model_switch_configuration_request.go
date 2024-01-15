package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SwitchConfigurationRequest struct {
	ConfigId string `json:"config_id"`

	Body *ApplyConfigurationRequestBody `json:"body,omitempty"`
}

func (o SwitchConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchConfigurationRequest struct{}"
	}

	return strings.Join([]string{"SwitchConfigurationRequest", string(data)}, " ")
}
