package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PrivateNameServer struct {
	Priority *int32 `json:"priority,omitempty"`

	Address *string `json:"address,omitempty"`
}

func (o PrivateNameServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateNameServer struct{}"
	}

	return strings.Join([]string{"PrivateNameServer", string(data)}, " ")
}
