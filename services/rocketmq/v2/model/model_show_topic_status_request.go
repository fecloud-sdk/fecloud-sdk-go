package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowTopicStatusRequest struct {
	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`
}

func (o ShowTopicStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowTopicStatusRequest", string(data)}, " ")
}
