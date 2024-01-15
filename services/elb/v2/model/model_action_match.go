package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ActionMatch struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o ActionMatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionMatch struct{}"
	}

	return strings.Join([]string{"ActionMatch", string(data)}, " ")
}
