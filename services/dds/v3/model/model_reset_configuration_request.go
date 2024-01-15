package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetConfigurationRequest struct {
	ConfigId string `json:"config_id"`
}

func (o ResetConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ResetConfigurationRequest", string(data)}, " ")
}
