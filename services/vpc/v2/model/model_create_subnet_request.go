package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSubnetRequest struct {
	Body *CreateSubnetRequestBody `json:"body,omitempty"`
}

func (o CreateSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetRequest struct{}"
	}

	return strings.Join([]string{"CreateSubnetRequest", string(data)}, " ")
}
