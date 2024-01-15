package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HostName struct {
}

func (o HostName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostName struct{}"
	}

	return strings.Join([]string{"HostName", string(data)}, " ")
}
