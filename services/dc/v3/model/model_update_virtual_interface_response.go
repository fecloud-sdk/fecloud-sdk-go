package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateVirtualInterfaceResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	VirtualInterface *VirtualInterface `json:"virtual_interface,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o UpdateVirtualInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualInterfaceResponse struct{}"
	}

	return strings.Join([]string{"UpdateVirtualInterfaceResponse", string(data)}, " ")
}
