package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateTopicAccessPolicyRequest struct {
	InstanceId string `json:"instance_id"`

	Body *UpdateTopicAccessPolicyReq `json:"body,omitempty"`
}

func (o UpdateTopicAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyRequest", string(data)}, " ")
}
