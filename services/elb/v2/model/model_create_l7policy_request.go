package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateL7policyRequest struct {
	Body *CreateL7policyRequestBody `json:"body,omitempty"`
}

func (o CreateL7policyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7policyRequest struct{}"
	}

	return strings.Join([]string{"CreateL7policyRequest", string(data)}, " ")
}
