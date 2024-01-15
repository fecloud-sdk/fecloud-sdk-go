package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NeutronCreateFloatingIpRequestBody 创建floatingip对象
type NeutronCreateFloatingIpRequestBody struct {
	Floatingip *CreateFloatingIpOption `json:"floatingip"`
}

func (o NeutronCreateFloatingIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFloatingIpRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateFloatingIpRequestBody", string(data)}, " ")
}
