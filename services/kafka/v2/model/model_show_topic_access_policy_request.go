package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowTopicAccessPolicyRequest struct {
	InstanceId string `json:"instance_id"`

	TopicName string `json:"topic_name"`
}

func (o ShowTopicAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowTopicAccessPolicyRequest", string(data)}, " ")
}
