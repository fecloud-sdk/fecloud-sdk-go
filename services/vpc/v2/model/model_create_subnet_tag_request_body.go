package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSubnetTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateSubnetTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSubnetTagRequestBody", string(data)}, " ")
}
