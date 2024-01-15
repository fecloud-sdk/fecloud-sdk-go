package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DiffConfigurationRequest struct {
	SourceConfigurationId string `json:"source_configuration_id"`

	TargetConfigurationId string `json:"target_configuration_id"`
}

func (o DiffConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiffConfigurationRequest struct{}"
	}

	return strings.Join([]string{"DiffConfigurationRequest", string(data)}, " ")
}
