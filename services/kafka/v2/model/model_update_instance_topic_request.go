package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceTopicRequest struct {
	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceTopicReq `json:"body,omitempty"`
}

func (o UpdateInstanceTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceTopicRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceTopicRequest", string(data)}, " ")
}
