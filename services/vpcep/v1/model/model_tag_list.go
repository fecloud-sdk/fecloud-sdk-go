package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TagList struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o TagList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagList struct{}"
	}

	return strings.Join([]string{"TagList", string(data)}, " ")
}
