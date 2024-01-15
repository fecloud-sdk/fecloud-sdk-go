package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateWhitelistReq struct {
	TenantId *string `json:"tenant_id,omitempty"`

	ListenerId string `json:"listener_id"`

	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`

	Whitelist *string `json:"whitelist,omitempty"`
}

func (o CreateWhitelistReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWhitelistReq struct{}"
	}

	return strings.Join([]string{"CreateWhitelistReq", string(data)}, " ")
}
