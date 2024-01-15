package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowKeyStoreRequest struct {
	KeystoreId string `json:"keystore_id"`
}

func (o ShowKeyStoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeyStoreRequest struct{}"
	}

	return strings.Join([]string{"ShowKeyStoreRequest", string(data)}, " ")
}
