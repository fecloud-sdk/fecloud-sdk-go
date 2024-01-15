package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListConsumeGroupAccessPolicyResponse struct {
	Policies *[]ListAccessPolicyRespPolicies `json:"policies,omitempty"`

	Total float32 `json:"total,omitempty"`

	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListConsumeGroupAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumeGroupAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListConsumeGroupAccessPolicyResponse", string(data)}, " ")
}
