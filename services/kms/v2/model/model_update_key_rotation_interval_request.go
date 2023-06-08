package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// Request Object
type UpdateKeyRotationIntervalRequest struct {
	Body *UpdateKeyRotationIntervalRequestBody `json:"body,omitempty"`
}

func (o UpdateKeyRotationIntervalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyRotationIntervalRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeyRotationIntervalRequest", string(data)}, " ")
}
