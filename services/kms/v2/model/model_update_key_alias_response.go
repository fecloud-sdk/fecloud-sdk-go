package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateKeyAliasResponse Response Object
type UpdateKeyAliasResponse struct {
	KeyInfo        *KeyAliasInfo `json:"key_info,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateKeyAliasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyAliasResponse struct{}"
	}

	return strings.Join([]string{"UpdateKeyAliasResponse", string(data)}, " ")
}
