package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EncryptDataResponse struct {
	KeyId *string `json:"key_id,omitempty"`

	CipherText     *string `json:"cipher_text,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EncryptDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDataResponse struct{}"
	}

	return strings.Join([]string{"EncryptDataResponse", string(data)}, " ")
}
