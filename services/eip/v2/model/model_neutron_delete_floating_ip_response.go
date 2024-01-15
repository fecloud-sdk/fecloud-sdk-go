package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NeutronDeleteFloatingIpResponse Response Object
type NeutronDeleteFloatingIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteFloatingIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteFloatingIpResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteFloatingIpResponse", string(data)}, " ")
}
