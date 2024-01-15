package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Metadata struct {
	CryptKeyId *string `json:"crypt_key_id,omitempty"`

	DedicatedFlavor *string `json:"dedicated_flavor,omitempty"`

	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`

	ExpandType *string `json:"expand_type,omitempty"`
}

func (o Metadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metadata struct{}"
	}

	return strings.Join([]string{"Metadata", string(data)}, " ")
}
