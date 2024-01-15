package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DefaultGroup struct {
}

func (o DefaultGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefaultGroup struct{}"
	}

	return strings.Join([]string{"DefaultGroup", string(data)}, " ")
}
