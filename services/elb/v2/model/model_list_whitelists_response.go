package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListWhitelistsResponse struct {
	Whitelists     *[]WhitelistResp `json:"whitelists,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListWhitelistsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWhitelistsResponse struct{}"
	}

	return strings.Join([]string{"ListWhitelistsResponse", string(data)}, " ")
}
