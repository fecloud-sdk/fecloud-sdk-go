package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteVirtualInterfaceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVirtualInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVirtualInterfaceResponse struct{}"
	}

	return strings.Join([]string{"DeleteVirtualInterfaceResponse", string(data)}, " ")
}
