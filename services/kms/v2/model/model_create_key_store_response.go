package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateKeyStoreResponse struct {
	Keystore       *KeystoreInfo `json:"keystore,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateKeyStoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeyStoreResponse struct{}"
	}

	return strings.Join([]string{"CreateKeyStoreResponse", string(data)}, " ")
}
