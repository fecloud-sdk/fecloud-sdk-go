package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HostNum struct {
}

func (o HostNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostNum struct{}"
	}

	return strings.Join([]string{"HostNum", string(data)}, " ")
}
