package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Ipv6Bandwidth struct {
	Id *string `json:"id,omitempty"`
}

func (o Ipv6Bandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Ipv6Bandwidth struct{}"
	}

	return strings.Join([]string{"Ipv6Bandwidth", string(data)}, " ")
}
