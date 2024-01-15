package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TagPlain struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o TagPlain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagPlain struct{}"
	}

	return strings.Join([]string{"TagPlain", string(data)}, " ")
}
