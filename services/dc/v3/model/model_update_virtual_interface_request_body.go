package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateVirtualInterfaceRequestBody struct {
	VirtualInterface *UpdateVirtualInterface `json:"virtual_interface"`
}

func (o UpdateVirtualInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVirtualInterfaceRequestBody", string(data)}, " ")
}
