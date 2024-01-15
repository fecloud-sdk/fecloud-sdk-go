package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteWhitelistRequest struct {
	WhitelistId string `json:"whitelist_id"`
}

func (o DeleteWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWhitelistRequest struct{}"
	}

	return strings.Join([]string{"DeleteWhitelistRequest", string(data)}, " ")
}
