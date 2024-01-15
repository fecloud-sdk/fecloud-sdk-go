package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceRemarkRequest struct {
	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceRemarkRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceRemarkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRemarkRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRemarkRequest", string(data)}, " ")
}
