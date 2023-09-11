package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DisassociateServerVirtualIpRequestBody This is a auto create Body Object
type DisassociateServerVirtualIpRequestBody struct {
	Nic *DisassociateServerVirtualIpOption `json:"nic"`
}

func (o DisassociateServerVirtualIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateServerVirtualIpRequestBody struct{}"
	}

	return strings.Join([]string{"DisassociateServerVirtualIpRequestBody", string(data)}, " ")
}
