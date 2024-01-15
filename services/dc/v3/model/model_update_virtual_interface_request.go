package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateVirtualInterfaceRequest struct {
	VirtualInterfaceId string `json:"virtual_interface_id"`

	Body *UpdateVirtualInterfaceRequestBody `json:"body,omitempty"`
}

func (o UpdateVirtualInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualInterfaceRequest struct{}"
	}

	return strings.Join([]string{"UpdateVirtualInterfaceRequest", string(data)}, " ")
}
