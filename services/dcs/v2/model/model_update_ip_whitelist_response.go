package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateIpWhitelistResponse Response Object
type UpdateIpWhitelistResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateIpWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpWhitelistResponse struct{}"
	}

	return strings.Join([]string{"UpdateIpWhitelistResponse", string(data)}, " ")
}
