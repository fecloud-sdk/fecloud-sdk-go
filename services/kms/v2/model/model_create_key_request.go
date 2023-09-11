package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateKeyRequest Request Object
type CreateKeyRequest struct {
	Body *CreateKeyRequestBody `json:"body,omitempty"`
}

func (o CreateKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeyRequest struct{}"
	}

	return strings.Join([]string{"CreateKeyRequest", string(data)}, " ")
}
