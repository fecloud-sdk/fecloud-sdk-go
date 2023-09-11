package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteKeyRequest Request Object
type DeleteKeyRequest struct {
	Body *ScheduleKeyDeletionRequestBody `json:"body,omitempty"`
}

func (o DeleteKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKeyRequest struct{}"
	}

	return strings.Join([]string{"DeleteKeyRequest", string(data)}, " ")
}
