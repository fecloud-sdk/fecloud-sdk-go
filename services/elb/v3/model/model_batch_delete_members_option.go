package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteMembersOption struct {
	Id *string `json:"id,omitempty"`

	Address *string `json:"address,omitempty"`

	ProtocolPort *int32 `json:"protocol_port,omitempty"`
}

func (o BatchDeleteMembersOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersOption struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersOption", string(data)}, " ")
}
