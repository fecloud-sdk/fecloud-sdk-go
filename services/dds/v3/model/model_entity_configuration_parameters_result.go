package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EntityConfigurationParametersResult struct {
	Name string `json:"name"`

	Value string `json:"value"`

	ValueRange string `json:"value_range"`

	RestartRequired bool `json:"restart_required"`

	Readonly bool `json:"readonly"`

	Type string `json:"type"`

	Description string `json:"description"`
}

func (o EntityConfigurationParametersResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EntityConfigurationParametersResult struct{}"
	}

	return strings.Join([]string{"EntityConfigurationParametersResult", string(data)}, " ")
}
