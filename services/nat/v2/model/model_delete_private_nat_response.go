package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeletePrivateNatResponse Response Object
type DeletePrivateNatResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateNatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateNatResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateNatResponse", string(data)}, " ")
}
