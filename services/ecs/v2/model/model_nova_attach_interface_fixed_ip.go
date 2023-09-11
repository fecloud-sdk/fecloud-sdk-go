package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NovaAttachInterfaceFixedIp
type NovaAttachInterfaceFixedIp struct {

	// IP地址。
	IpAddress *string `json:"ip_address,omitempty"`
}

func (o NovaAttachInterfaceFixedIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAttachInterfaceFixedIp struct{}"
	}

	return strings.Join([]string{"NovaAttachInterfaceFixedIp", string(data)}, " ")
}
