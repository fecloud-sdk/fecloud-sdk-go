package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateVpcResourceTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateVpcResourceTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcResourceTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVpcResourceTagRequestBody", string(data)}, " ")
}
