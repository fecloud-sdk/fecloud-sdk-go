package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteWhitelistResponse Response Object
type DeleteWhitelistResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWhitelistResponse struct{}"
	}

	return strings.Join([]string{"DeleteWhitelistResponse", string(data)}, " ")
}
