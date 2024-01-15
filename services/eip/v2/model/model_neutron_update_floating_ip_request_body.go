package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NeutronUpdateFloatingIpRequestBody 更新floatingip的请求体
type NeutronUpdateFloatingIpRequestBody struct {
	Floatingip *UpdateFloatingIpOption `json:"floatingip"`
}

func (o NeutronUpdateFloatingIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFloatingIpRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFloatingIpRequestBody", string(data)}, " ")
}
