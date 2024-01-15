package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Description struct {
}

func (o Description) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Description struct{}"
	}

	return strings.Join([]string{"Description", string(data)}, " ")
}
