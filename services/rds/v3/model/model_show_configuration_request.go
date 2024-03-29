package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConfigurationRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	ConfigId string `json:"config_id"`
}

func (o ShowConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationRequest", string(data)}, " ")
}
