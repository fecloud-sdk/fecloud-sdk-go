package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateWhitelistRequestBody struct {
	Whitelist *CreateWhitelistReq `json:"whitelist"`
}

func (o CreateWhitelistRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWhitelistRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWhitelistRequestBody", string(data)}, " ")
}
