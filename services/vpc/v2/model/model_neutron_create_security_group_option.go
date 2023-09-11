package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateSecurityGroupOption struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`
}

func (o NeutronCreateSecurityGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupOption struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupOption", string(data)}, " ")
}
