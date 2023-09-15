package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// SetRecyclePolicyRequest Request Object
type SetRecyclePolicyRequest struct {
	Body *RecyclePolicyRequestBody `json:"body,omitempty"`
}

func (o SetRecyclePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRecyclePolicyRequest struct{}"
	}

	return strings.Join([]string{"SetRecyclePolicyRequest", string(data)}, " ")
}
