package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateVpcOption struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o UpdateVpcOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcOption struct{}"
	}

	return strings.Join([]string{"UpdateVpcOption", string(data)}, " ")
}
