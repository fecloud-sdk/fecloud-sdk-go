package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeShareNameReq struct {
	ChangeName *ShareName `json:"change_name"`
}

func (o ChangeShareNameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeShareNameReq struct{}"
	}

	return strings.Join([]string{"ChangeShareNameReq", string(data)}, " ")
}
