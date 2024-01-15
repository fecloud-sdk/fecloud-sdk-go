package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteTopicRequest struct {
	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`
}

func (o DeleteTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTopicRequest struct{}"
	}

	return strings.Join([]string{"DeleteTopicRequest", string(data)}, " ")
}
