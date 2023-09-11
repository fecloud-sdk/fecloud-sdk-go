package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// KeyDescriptionInfo 密钥描述信息。
type KeyDescriptionInfo struct {

	// 密钥ID。
	KeyId *string `json:"key_id,omitempty"`

	// 密钥描述。
	KeyDescription *string `json:"key_description,omitempty"`
}

func (o KeyDescriptionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyDescriptionInfo struct{}"
	}

	return strings.Join([]string{"KeyDescriptionInfo", string(data)}, " ")
}
