package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateWhitelistRequestBody struct {
	Whitelist *UpdateWhitelistReq `json:"whitelist"`
}

func (o UpdateWhitelistRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhitelistRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateWhitelistRequestBody", string(data)}, " ")
}
