package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateServerOption struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Hostname *string `json:"hostname,omitempty"`
}

func (o UpdateServerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerOption struct{}"
	}

	return strings.Join([]string{"UpdateServerOption", string(data)}, " ")
}
