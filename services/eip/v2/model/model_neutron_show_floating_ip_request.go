package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NeutronShowFloatingIpRequest Request Object
type NeutronShowFloatingIpRequest struct {

	// floatingip的ID
	FloatingipId string `json:"floatingip_id"`
}

func (o NeutronShowFloatingIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowFloatingIpRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowFloatingIpRequest", string(data)}, " ")
}
