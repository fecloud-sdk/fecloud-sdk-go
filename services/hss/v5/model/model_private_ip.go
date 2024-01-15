package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PrivateIp struct {
}

func (o PrivateIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateIp struct{}"
	}

	return strings.Join([]string{"PrivateIp", string(data)}, " ")
}
