package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePrivateNatRequestBody struct {
	Gateway *CreatePrivateNatOption `json:"gateway"`
}

func (o CreatePrivateNatRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateNatRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePrivateNatRequestBody", string(data)}, " ")
}
