package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QueryResourceTagItem struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o QueryResourceTagItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryResourceTagItem struct{}"
	}

	return strings.Join([]string{"QueryResourceTagItem", string(data)}, " ")
}
