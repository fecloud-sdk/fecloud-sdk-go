package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateVpcRequestBody struct {
	Vpc *UpdateVpcOption `json:"vpc"`
}

func (o UpdateVpcRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVpcRequestBody", string(data)}, " ")
}
