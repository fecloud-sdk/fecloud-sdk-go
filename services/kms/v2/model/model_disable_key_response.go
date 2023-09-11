package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DisableKeyResponse Response Object
type DisableKeyResponse struct {
	KeyInfo        *KeyStatusInfo `json:"key_info,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o DisableKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableKeyResponse struct{}"
	}

	return strings.Join([]string{"DisableKeyResponse", string(data)}, " ")
}
