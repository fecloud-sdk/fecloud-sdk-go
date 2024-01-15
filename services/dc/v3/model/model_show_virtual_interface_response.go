package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowVirtualInterfaceResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	VirtualInterface *VirtualInterface `json:"virtual_interface,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o ShowVirtualInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVirtualInterfaceResponse struct{}"
	}

	return strings.Join([]string{"ShowVirtualInterfaceResponse", string(data)}, " ")
}
