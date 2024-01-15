package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetUserPasswrodResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetUserPasswrodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetUserPasswrodResponse struct{}"
	}

	return strings.Join([]string{"ResetUserPasswrodResponse", string(data)}, " ")
}
