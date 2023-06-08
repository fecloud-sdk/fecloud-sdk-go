package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// Request Object
type EncryptDataRequest struct {
	Body *EncryptDataRequestBody `json:"body,omitempty"`
}

func (o EncryptDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDataRequest struct{}"
	}

	return strings.Join([]string{"EncryptDataRequest", string(data)}, " ")
}
