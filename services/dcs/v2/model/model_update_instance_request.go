package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceRequest struct {
	InstanceId string `json:"instance_id"`

	Body *ModifyInstanceBody `json:"body,omitempty"`
}

func (o UpdateInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRequest", string(data)}, " ")
}
