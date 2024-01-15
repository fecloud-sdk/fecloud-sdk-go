package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type KeKInfo struct {
	KeyId *string `json:"key_id,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`
}

func (o KeKInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeKInfo struct{}"
	}

	return strings.Join([]string{"KeKInfo", string(data)}, " ")
}
