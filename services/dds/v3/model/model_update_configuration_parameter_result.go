package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateConfigurationParameterResult struct {
	EntityId string `json:"entity_id"`

	ParameterValues map[string]string `json:"parameter_values"`
}

func (o UpdateConfigurationParameterResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationParameterResult struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationParameterResult", string(data)}, " ")
}
