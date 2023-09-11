package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// AssociateServerVirtualIpRequest Request Object
type AssociateServerVirtualIpRequest struct {

	// 云服务器网卡ID。
	NicId string `json:"nic_id"`

	Body *AssociateServerVirtualIpRequestBody `json:"body,omitempty"`
}

func (o AssociateServerVirtualIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateServerVirtualIpRequest struct{}"
	}

	return strings.Join([]string{"AssociateServerVirtualIpRequest", string(data)}, " ")
}
