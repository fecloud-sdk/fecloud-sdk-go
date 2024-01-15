package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateConfigurationParameterRequestBody struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	ParameterValues map[string]string `json:"parameter_values,omitempty"`
}

func (o UpdateConfigurationParameterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationParameterRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationParameterRequestBody", string(data)}, " ")
}
