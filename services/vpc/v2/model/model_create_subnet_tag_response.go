package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSubnetTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSubnetTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetTagResponse struct{}"
	}

	return strings.Join([]string{"CreateSubnetTagResponse", string(data)}, " ")
}
