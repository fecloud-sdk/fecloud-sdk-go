package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type WhitelistResp struct {
	Id string `json:"id"`

	TenantId string `json:"tenant_id"`

	ListenerId string `json:"listener_id"`

	EnableWhitelist bool `json:"enable_whitelist"`

	Whitelist string `json:"whitelist"`
}

func (o WhitelistResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WhitelistResp struct{}"
	}

	return strings.Join([]string{"WhitelistResp", string(data)}, " ")
}
