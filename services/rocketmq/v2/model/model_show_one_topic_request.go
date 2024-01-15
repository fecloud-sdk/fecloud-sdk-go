package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowOneTopicRequest struct {
	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`
}

func (o ShowOneTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOneTopicRequest struct{}"
	}

	return strings.Join([]string{"ShowOneTopicRequest", string(data)}, " ")
}
