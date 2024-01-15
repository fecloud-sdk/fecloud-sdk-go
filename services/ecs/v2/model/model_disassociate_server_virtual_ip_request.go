package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DisassociateServerVirtualIpRequest struct {
	NicId string `json:"nic_id"`

	Body *DisassociateServerVirtualIpRequestBody `json:"body,omitempty"`
}

func (o DisassociateServerVirtualIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateServerVirtualIpRequest struct{}"
	}

	return strings.Join([]string{"DisassociateServerVirtualIpRequest", string(data)}, " ")
}
