package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateVirtualInterfaceRequest struct {
	Body *CreateVirtualInterfaceRequestBody `json:"body,omitempty"`
}

func (o CreateVirtualInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualInterfaceRequest struct{}"
	}

	return strings.Join([]string{"CreateVirtualInterfaceRequest", string(data)}, " ")
}
