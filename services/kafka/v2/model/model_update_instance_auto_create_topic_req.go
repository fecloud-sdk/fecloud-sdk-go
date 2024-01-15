package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceAutoCreateTopicReq struct {
	EnableAutoTopic bool `json:"enable_auto_topic"`
}

func (o UpdateInstanceAutoCreateTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceAutoCreateTopicReq struct{}"
	}

	return strings.Join([]string{"UpdateInstanceAutoCreateTopicReq", string(data)}, " ")
}
