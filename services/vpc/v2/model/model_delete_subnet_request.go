package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteSubnetRequest struct {
	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`
}

func (o DeleteSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubnetRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubnetRequest", string(data)}, " ")
}
