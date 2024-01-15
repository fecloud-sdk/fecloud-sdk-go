package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListTopicAccessPolicyResponse struct {
	Policies *[]ListAccessPolicyRespPolicies `json:"policies,omitempty"`

	Total float32 `json:"total,omitempty"`

	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTopicAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListTopicAccessPolicyResponse", string(data)}, " ")
}
