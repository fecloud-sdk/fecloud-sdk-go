package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowInstanceTopicDetailRequest struct {
	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`
}

func (o ShowInstanceTopicDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTopicDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceTopicDetailRequest", string(data)}, " ")
}
