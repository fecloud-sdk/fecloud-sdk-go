package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteVirtualInterfaceRequest struct {
	VirtualInterfaceId string `json:"virtual_interface_id"`
}

func (o DeleteVirtualInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVirtualInterfaceRequest struct{}"
	}

	return strings.Join([]string{"DeleteVirtualInterfaceRequest", string(data)}, " ")
}
