package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateConfigurationsRequest struct {
	InstanceId string `json:"instance_id"`

	Body *ModifyRedisConfigBody `json:"body,omitempty"`
}

func (o UpdateConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationsRequest", string(data)}, " ")
}
