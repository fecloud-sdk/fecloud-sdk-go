package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteVpcRequest struct {
	VpcId string `json:"vpc_id"`
}

func (o DeleteVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpcRequest", string(data)}, " ")
}
