package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateL7policyRequestBody struct {
	L7policy *CreateL7policyReq `json:"l7policy"`
}

func (o CreateL7policyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7policyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateL7policyRequestBody", string(data)}, " ")
}
