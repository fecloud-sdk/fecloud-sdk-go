package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Unit struct {
}

func (o Unit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Unit struct{}"
	}

	return strings.Join([]string{"Unit", string(data)}, " ")
}
