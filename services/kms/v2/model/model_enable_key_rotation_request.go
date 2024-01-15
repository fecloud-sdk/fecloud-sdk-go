package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EnableKeyRotationRequest struct {
	Body *OperateKeyRequestBody `json:"body,omitempty"`
}

func (o EnableKeyRotationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableKeyRotationRequest struct{}"
	}

	return strings.Join([]string{"EnableKeyRotationRequest", string(data)}, " ")
}
