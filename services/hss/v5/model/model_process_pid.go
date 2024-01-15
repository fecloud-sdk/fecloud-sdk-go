package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ProcessPid struct {
}

func (o ProcessPid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessPid struct{}"
	}

	return strings.Join([]string{"ProcessPid", string(data)}, " ")
}
