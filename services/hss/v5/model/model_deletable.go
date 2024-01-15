package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Deletable struct {
}

func (o Deletable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Deletable struct{}"
	}

	return strings.Join([]string{"Deletable", string(data)}, " ")
}
