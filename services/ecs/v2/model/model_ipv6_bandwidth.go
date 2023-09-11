package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// Ipv6Bandwidth
type Ipv6Bandwidth struct {

	// IPv6带宽的ID。
	Id *string `json:"id,omitempty"`
}

func (o Ipv6Bandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Ipv6Bandwidth struct{}"
	}

	return strings.Join([]string{"Ipv6Bandwidth", string(data)}, " ")
}
