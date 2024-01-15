package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetManagerPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetManagerPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetManagerPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetManagerPasswordResponse", string(data)}, " ")
}
