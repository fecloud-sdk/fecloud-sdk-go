package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowTopicAccessPolicyResponse struct {
	Name *string `json:"name,omitempty"`

	TopicType *int32 `json:"topic_type,omitempty"`

	Policies       *[]PolicyEntity `json:"policies,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowTopicAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowTopicAccessPolicyResponse", string(data)}, " ")
}
