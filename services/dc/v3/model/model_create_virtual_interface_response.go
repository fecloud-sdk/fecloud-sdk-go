package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateVirtualInterfaceResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	VirtualInterface *VirtualInterface `json:"virtual_interface,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o CreateVirtualInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualInterfaceResponse struct{}"
	}

	return strings.Join([]string{"CreateVirtualInterfaceResponse", string(data)}, " ")
}
