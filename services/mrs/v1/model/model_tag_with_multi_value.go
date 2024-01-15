package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TagWithMultiValue struct {
	Key string `json:"key"`

	Values *[]string `json:"values,omitempty"`
}

func (o TagWithMultiValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagWithMultiValue struct{}"
	}

	return strings.Join([]string{"TagWithMultiValue", string(data)}, " ")
}
