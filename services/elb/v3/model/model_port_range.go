package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PortRange struct {
	StartPort *int32 `json:"start_port,omitempty"`

	EndPort *int32 `json:"end_port,omitempty"`
}

func (o PortRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortRange struct{}"
	}

	return strings.Join([]string{"PortRange", string(data)}, " ")
}
