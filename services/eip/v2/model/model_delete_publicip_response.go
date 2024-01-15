package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeletePublicipResponse Response Object
type DeletePublicipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePublicipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublicipResponse struct{}"
	}

	return strings.Join([]string{"DeletePublicipResponse", string(data)}, " ")
}
