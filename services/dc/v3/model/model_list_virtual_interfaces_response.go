package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListVirtualInterfacesResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	VirtualInterfaces *[]VirtualInterface `json:"virtual_interfaces,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListVirtualInterfacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVirtualInterfacesResponse struct{}"
	}

	return strings.Join([]string{"ListVirtualInterfacesResponse", string(data)}, " ")
}
