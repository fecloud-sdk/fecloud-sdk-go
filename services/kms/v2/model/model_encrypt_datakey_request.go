package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EncryptDatakeyRequest struct {
	Body *EncryptDatakeyRequestBody `json:"body,omitempty"`
}

func (o EncryptDatakeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDatakeyRequest struct{}"
	}

	return strings.Join([]string{"EncryptDatakeyRequest", string(data)}, " ")
}
