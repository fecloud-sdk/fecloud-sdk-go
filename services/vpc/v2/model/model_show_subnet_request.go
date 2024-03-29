package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSubnetRequest struct {
	SubnetId string `json:"subnet_id"`
}

func (o ShowSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubnetRequest struct{}"
	}

	return strings.Join([]string{"ShowSubnetRequest", string(data)}, " ")
}
