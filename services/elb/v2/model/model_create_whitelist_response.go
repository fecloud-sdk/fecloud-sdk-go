package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateWhitelistResponse struct {
	Whitelist      *WhitelistResp `json:"whitelist,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWhitelistResponse struct{}"
	}

	return strings.Join([]string{"CreateWhitelistResponse", string(data)}, " ")
}
