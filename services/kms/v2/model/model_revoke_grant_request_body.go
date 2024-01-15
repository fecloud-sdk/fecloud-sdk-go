package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RevokeGrantRequestBody struct {
	KeyId string `json:"key_id"`

	GrantId string `json:"grant_id"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o RevokeGrantRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeGrantRequestBody struct{}"
	}

	return strings.Join([]string{"RevokeGrantRequestBody", string(data)}, " ")
}
