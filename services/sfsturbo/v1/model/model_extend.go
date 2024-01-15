package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Extend struct {
	NewSize int32 `json:"new_size"`
}

func (o Extend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Extend struct{}"
	}

	return strings.Join([]string{"Extend", string(data)}, " ")
}
