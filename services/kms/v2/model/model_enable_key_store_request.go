package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EnableKeyStoreRequest struct {
	KeystoreId string `json:"keystore_id"`
}

func (o EnableKeyStoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableKeyStoreRequest struct{}"
	}

	return strings.Join([]string{"EnableKeyStoreRequest", string(data)}, " ")
}
