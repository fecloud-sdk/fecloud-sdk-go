package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TagValuesList struct {
	Key string `json:"key"`

	Values []string `json:"values"`
}

func (o TagValuesList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagValuesList struct{}"
	}

	return strings.Join([]string{"TagValuesList", string(data)}, " ")
}
