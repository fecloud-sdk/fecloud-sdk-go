package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ActionTag struct {
	Key string `json:"key"`

	Values []string `json:"values"`
}

func (o ActionTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionTag struct{}"
	}

	return strings.Join([]string{"ActionTag", string(data)}, " ")
}
