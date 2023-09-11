package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CancelKeyDeletionRequest Request Object
type CancelKeyDeletionRequest struct {
	Body *OperateKeyRequestBody `json:"body,omitempty"`
}

func (o CancelKeyDeletionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelKeyDeletionRequest struct{}"
	}

	return strings.Join([]string{"CancelKeyDeletionRequest", string(data)}, " ")
}
