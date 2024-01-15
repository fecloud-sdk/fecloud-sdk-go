package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowVirtualInterfaceRequest struct {
	Fields *[]string `json:"fields,omitempty"`

	VirtualInterfaceId string `json:"virtual_interface_id"`
}

func (o ShowVirtualInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVirtualInterfaceRequest struct{}"
	}

	return strings.Join([]string{"ShowVirtualInterfaceRequest", string(data)}, " ")
}
