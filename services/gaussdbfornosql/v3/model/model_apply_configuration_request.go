package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ApplyConfigurationRequest struct {
	ConfigId string `json:"config_id"`

	Body *ApplyConfigurationRequestBody `json:"body,omitempty"`
}

func (o ApplyConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationRequest", string(data)}, " ")
}
