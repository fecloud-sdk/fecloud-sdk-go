package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EnableKeyStoreResponse struct {
	Keystore       *KeyStoreStateInfo `json:"keystore,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o EnableKeyStoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableKeyStoreResponse struct{}"
	}

	return strings.Join([]string{"EnableKeyStoreResponse", string(data)}, " ")
}
