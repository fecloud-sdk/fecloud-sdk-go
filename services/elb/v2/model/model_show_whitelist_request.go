package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowWhitelistRequest struct {
	WhitelistId string `json:"whitelist_id"`
}

func (o ShowWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWhitelistRequest struct{}"
	}

	return strings.Join([]string{"ShowWhitelistRequest", string(data)}, " ")
}
