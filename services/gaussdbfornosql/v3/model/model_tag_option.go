package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TagOption struct {
	Key string `json:"key"`

	Values []string `json:"values"`
}

func (o TagOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagOption struct{}"
	}

	return strings.Join([]string{"TagOption", string(data)}, " ")
}
