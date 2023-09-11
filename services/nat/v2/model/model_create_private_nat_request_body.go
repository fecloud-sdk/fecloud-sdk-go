package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreatePrivateNatRequestBody 创建私网NAT网关实例的请求体。
type CreatePrivateNatRequestBody struct {
	Gateway *CreatePrivateNatOption `json:"gateway"`
}

func (o CreatePrivateNatRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateNatRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePrivateNatRequestBody", string(data)}, " ")
}
