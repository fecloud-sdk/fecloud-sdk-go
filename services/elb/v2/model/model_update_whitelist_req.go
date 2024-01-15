package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateWhitelistReq struct {
	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`

	Whitelist *string `json:"whitelist,omitempty"`
}

func (o UpdateWhitelistReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhitelistReq struct{}"
	}

	return strings.Join([]string{"UpdateWhitelistReq", string(data)}, " ")
}
