package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// EnableKeyStoreRequest Request Object
type EnableKeyStoreRequest struct {

	// 密钥库ID
	KeystoreId string `json:"keystore_id"`
}

func (o EnableKeyStoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableKeyStoreRequest struct{}"
	}

	return strings.Join([]string{"EnableKeyStoreRequest", string(data)}, " ")
}
