package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateVpcRequest struct {
	VpcId string `json:"vpc_id"`

	Body *UpdateVpcRequestBody `json:"body,omitempty"`
}

func (o UpdateVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcRequest struct{}"
	}

	return strings.Join([]string{"UpdateVpcRequest", string(data)}, " ")
}
