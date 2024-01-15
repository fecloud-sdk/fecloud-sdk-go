package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConfigurationParameterRequest struct {
	ConfigId string `json:"config_id"`
}

func (o ShowConfigurationParameterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationParameterRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationParameterRequest", string(data)}, " ")
}
