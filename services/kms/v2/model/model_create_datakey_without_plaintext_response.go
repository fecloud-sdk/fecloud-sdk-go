package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateDatakeyWithoutPlaintextResponse struct {
	KeyId *string `json:"key_id,omitempty"`

	CipherText     *string `json:"cipher_text,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDatakeyWithoutPlaintextResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatakeyWithoutPlaintextResponse struct{}"
	}

	return strings.Join([]string{"CreateDatakeyWithoutPlaintextResponse", string(data)}, " ")
}
