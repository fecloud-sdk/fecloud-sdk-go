package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateConfigurationParameterRequest struct {
	ConfigId string `json:"config_id"`

	Body *UpdateConfigurationParameterRequestBody `json:"body,omitempty"`
}

func (o UpdateConfigurationParameterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationParameterRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationParameterRequest", string(data)}, " ")
}
