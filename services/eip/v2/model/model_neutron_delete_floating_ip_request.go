package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NeutronDeleteFloatingIpRequest Request Object
type NeutronDeleteFloatingIpRequest struct {

	// floatingip的ID
	FloatingipId string `json:"floatingip_id"`
}

func (o NeutronDeleteFloatingIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteFloatingIpRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteFloatingIpRequest", string(data)}, " ")
}
