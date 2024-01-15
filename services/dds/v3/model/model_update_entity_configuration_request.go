package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEntityConfigurationRequest struct {
	InstanceId string `json:"instance_id"`

	Body *UpdateConfigurationParameterResult `json:"body,omitempty"`
}

func (o UpdateEntityConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEntityConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpdateEntityConfigurationRequest", string(data)}, " ")
}
