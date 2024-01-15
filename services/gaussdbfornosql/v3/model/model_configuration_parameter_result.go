package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ConfigurationParameterResult struct {
	Name string `json:"name"`

	Value string `json:"value"`

	RestartRequired bool `json:"restart_required"`

	Readonly bool `json:"readonly"`

	ValueRange string `json:"value_range"`

	Type string `json:"type"`

	Description string `json:"description"`
}

func (o ConfigurationParameterResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationParameterResult struct{}"
	}

	return strings.Join([]string{"ConfigurationParameterResult", string(data)}, " ")
}
