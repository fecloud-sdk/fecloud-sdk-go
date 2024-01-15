package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceAutoCreateTopicRequest struct {
	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceAutoCreateTopicReq `json:"body,omitempty"`
}

func (o UpdateInstanceAutoCreateTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceAutoCreateTopicRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceAutoCreateTopicRequest", string(data)}, " ")
}
