package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateConfigurationRequest struct {
	ConfigId string `json:"config_id"`

	Body *UpdateConfigurationRequestBody `json:"body,omitempty"`
}

func (o UpdateConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationRequest", string(data)}, " ")
}
