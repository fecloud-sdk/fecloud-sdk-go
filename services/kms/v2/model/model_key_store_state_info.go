package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// 密钥库状态详情
type KeyStoreStateInfo struct {

	// 密钥库ID
	KeystoreId *string `json:"keystore_id,omitempty"`

	// 密钥库状态
	KeystoreState *string `json:"keystore_state,omitempty"`
}

func (o KeyStoreStateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyStoreStateInfo struct{}"
	}

	return strings.Join([]string{"KeyStoreStateInfo", string(data)}, " ")
}
