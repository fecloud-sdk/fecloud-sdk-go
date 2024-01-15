package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HostId struct {
}

func (o HostId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostId struct{}"
	}

	return strings.Join([]string{"HostId", string(data)}, " ")
}
