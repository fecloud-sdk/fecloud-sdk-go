package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type LoginUserName struct {
}

func (o LoginUserName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginUserName struct{}"
	}

	return strings.Join([]string{"LoginUserName", string(data)}, " ")
}
