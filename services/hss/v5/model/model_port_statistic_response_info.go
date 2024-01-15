package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PortStatisticResponseInfo struct {
	Port *int32 `json:"port,omitempty"`

	Type *string `json:"type,omitempty"`

	Num *int32 `json:"num,omitempty"`
}

func (o PortStatisticResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortStatisticResponseInfo struct{}"
	}

	return strings.Join([]string{"PortStatisticResponseInfo", string(data)}, " ")
}
