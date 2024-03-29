package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DisableKeyRotationRequest struct {
	Body *OperateKeyRequestBody `json:"body,omitempty"`
}

func (o DisableKeyRotationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableKeyRotationRequest struct{}"
	}

	return strings.Join([]string{"DisableKeyRotationRequest", string(data)}, " ")
}
