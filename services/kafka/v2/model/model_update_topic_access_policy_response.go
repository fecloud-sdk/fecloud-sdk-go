package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateTopicAccessPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTopicAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyResponse", string(data)}, " ")
}
