package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateVirtualInterfaceRequestBody struct {
	VirtualInterface *CreateVirtualInterface `json:"virtual_interface"`
}

func (o CreateVirtualInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVirtualInterfaceRequestBody", string(data)}, " ")
}
