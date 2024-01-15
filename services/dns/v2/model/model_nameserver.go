package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Nameserver struct {
	Hostname *string `json:"hostname,omitempty"`

	Priority *int32 `json:"priority,omitempty"`
}

func (o Nameserver) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nameserver struct{}"
	}

	return strings.Join([]string{"Nameserver", string(data)}, " ")
}
