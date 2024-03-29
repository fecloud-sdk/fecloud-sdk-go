package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowWhitelistResponse struct {
	Whitelist      *WhitelistResp `json:"whitelist,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWhitelistResponse struct{}"
	}

	return strings.Join([]string{"ShowWhitelistResponse", string(data)}, " ")
}
