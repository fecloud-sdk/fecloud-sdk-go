package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateMembersOption struct {
	Id string `json:"id"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Name *string `json:"name,omitempty"`

	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	Weight *int32 `json:"weight,omitempty"`
}

func (o BatchUpdateMembersOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateMembersOption struct{}"
	}

	return strings.Join([]string{"BatchUpdateMembersOption", string(data)}, " ")
}
