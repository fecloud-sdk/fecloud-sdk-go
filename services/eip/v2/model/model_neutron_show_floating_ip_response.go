package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NeutronShowFloatingIpResponse Response Object
type NeutronShowFloatingIpResponse struct {
	Floatingip     *FloatingIpResp `json:"floatingip,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o NeutronShowFloatingIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowFloatingIpResponse struct{}"
	}

	return strings.Join([]string{"NeutronShowFloatingIpResponse", string(data)}, " ")
}
