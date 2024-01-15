package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ConfigurationParametersResult struct {
	Name string `json:"name"`

	Value string `json:"value"`

	Description string `json:"description"`

	Type string `json:"type"`

	ValueRange string `json:"value_range"`

	RestartRequired bool `json:"restart_required"`

	Readonly bool `json:"readonly"`
}

func (o ConfigurationParametersResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationParametersResult struct{}"
	}

	return strings.Join([]string{"ConfigurationParametersResult", string(data)}, " ")
}
