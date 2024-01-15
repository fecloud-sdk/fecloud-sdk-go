package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QueryTagItem struct {
	Key string `json:"key"`

	Values []string `json:"values"`
}

func (o QueryTagItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTagItem struct{}"
	}

	return strings.Join([]string{"QueryTagItem", string(data)}, " ")
}
