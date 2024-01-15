package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceNameRequest struct {
	InstanceId string `json:"instance_id"`

	Body *UpdateNameRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceNameRequest", string(data)}, " ")
}
