package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateVpcResponse struct {
	Vpc            *Vpc `json:"vpc,omitempty"`
	HttpStatusCode int  `json:"-"`
}

func (o UpdateVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcResponse struct{}"
	}

	return strings.Join([]string{"UpdateVpcResponse", string(data)}, " ")
}
