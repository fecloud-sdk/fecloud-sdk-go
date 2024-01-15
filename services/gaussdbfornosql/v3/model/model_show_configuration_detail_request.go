package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConfigurationDetailRequest struct {
	ConfigId string `json:"config_id"`
}

func (o ShowConfigurationDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationDetailRequest", string(data)}, " ")
}
