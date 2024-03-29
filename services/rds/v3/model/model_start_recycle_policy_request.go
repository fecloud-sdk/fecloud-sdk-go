package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type StartRecyclePolicyRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	Body *RecyclePolicyRequestBody `json:"body,omitempty"`
}

func (o StartRecyclePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartRecyclePolicyRequest struct{}"
	}

	return strings.Join([]string{"StartRecyclePolicyRequest", string(data)}, " ")
}
