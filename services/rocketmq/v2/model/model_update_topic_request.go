package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateTopicRequest struct {
	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`

	Body *UpdateTopicReq `json:"body,omitempty"`
}

func (o UpdateTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicRequest struct{}"
	}

	return strings.Join([]string{"UpdateTopicRequest", string(data)}, " ")
}
