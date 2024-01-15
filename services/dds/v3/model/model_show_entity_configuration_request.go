package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowEntityConfigurationRequest struct {
	InstanceId string `json:"instance_id"`

	EntityId string `json:"entity_id"`
}

func (o ShowEntityConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEntityConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ShowEntityConfigurationRequest", string(data)}, " ")
}
