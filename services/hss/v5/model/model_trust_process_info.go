package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TrustProcessInfo struct {
	Path *string `json:"path,omitempty"`

	Hash *string `json:"hash,omitempty"`
}

func (o TrustProcessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrustProcessInfo struct{}"
	}

	return strings.Join([]string{"TrustProcessInfo", string(data)}, " ")
}
