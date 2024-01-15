package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListTopicAccessPolicyRequest struct {
	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`

	Offset *string `json:"offset,omitempty"`

	Limit *string `json:"limit,omitempty"`
}

func (o ListTopicAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListTopicAccessPolicyRequest", string(data)}, " ")
}
