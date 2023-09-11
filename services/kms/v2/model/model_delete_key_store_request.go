package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteKeyStoreRequest Request Object
type DeleteKeyStoreRequest struct {

	// 密钥库ID
	KeystoreId string `json:"keystore_id"`
}

func (o DeleteKeyStoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKeyStoreRequest struct{}"
	}

	return strings.Join([]string{"DeleteKeyStoreRequest", string(data)}, " ")
}
