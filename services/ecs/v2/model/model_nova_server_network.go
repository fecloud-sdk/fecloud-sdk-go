package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NovaServerNetwork struct {
	Port *string `json:"port,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	FixedIp *string `json:"fixed_ip,omitempty"`
}

func (o NovaServerNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerNetwork struct{}"
	}

	return strings.Join([]string{"NovaServerNetwork", string(data)}, " ")
}
