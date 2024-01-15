package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeShareNameResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeShareNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeShareNameResponse struct{}"
	}

	return strings.Join([]string{"ChangeShareNameResponse", string(data)}, " ")
}
