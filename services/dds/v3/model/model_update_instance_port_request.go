package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstancePortRequest struct {
	InstanceId string `json:"instance_id"`

	Body *UpdatePortRequestBody `json:"body,omitempty"`
}

func (o UpdateInstancePortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstancePortRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstancePortRequest", string(data)}, " ")
}
