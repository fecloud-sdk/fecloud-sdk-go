package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PortResponseInfo struct {
	HostId *string `json:"host_id,omitempty"`

	Laddr *string `json:"laddr,omitempty"`

	Status *string `json:"status,omitempty"`

	Port *int32 `json:"port,omitempty"`

	Type *string `json:"type,omitempty"`

	Pid *int32 `json:"pid,omitempty"`

	Path *string `json:"path,omitempty"`
}

func (o PortResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortResponseInfo struct{}"
	}

	return strings.Join([]string{"PortResponseInfo", string(data)}, " ")
}
