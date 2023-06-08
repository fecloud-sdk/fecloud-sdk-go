package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// Request Object
type UpdateKeyAliasRequest struct {
	Body *UpdateKeyAliasRequestBody `json:"body,omitempty"`
}

func (o UpdateKeyAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyAliasRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeyAliasRequest", string(data)}, " ")
}
