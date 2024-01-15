package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type InstanceItemTagItem struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o InstanceItemTagItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceItemTagItem struct{}"
	}

	return strings.Join([]string{"InstanceItemTagItem", string(data)}, " ")
}
