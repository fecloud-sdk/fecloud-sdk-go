package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateWhitelistRequest struct {
	WhitelistId string `json:"whitelist_id"`

	Body *UpdateWhitelistRequestBody `json:"body,omitempty"`
}

func (o UpdateWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhitelistRequest struct{}"
	}

	return strings.Join([]string{"UpdateWhitelistRequest", string(data)}, " ")
}
