package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Value struct {
}

func (o Value) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Value struct{}"
	}

	return strings.Join([]string{"Value", string(data)}, " ")
}
