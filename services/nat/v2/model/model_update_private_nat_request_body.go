package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePrivateNatRequestBody struct {
	Gateway *UpdatePrivateNatOption `json:"gateway"`
}

func (o UpdatePrivateNatRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateNatRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePrivateNatRequestBody", string(data)}, " ")
}
