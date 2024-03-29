package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSubnetTagRequest struct {
	SubnetId string `json:"subnet_id"`

	Body *CreateSubnetTagRequestBody `json:"body,omitempty"`
}

func (o CreateSubnetTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetTagRequest struct{}"
	}

	return strings.Join([]string{"CreateSubnetTagRequest", string(data)}, " ")
}
