package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type LoginIp struct {
}

func (o LoginIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginIp struct{}"
	}

	return strings.Join([]string{"LoginIp", string(data)}, " ")
}
