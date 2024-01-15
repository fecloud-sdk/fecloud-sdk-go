package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSharedTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateSharedTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSharedTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSharedTagRequestBody", string(data)}, " ")
}
