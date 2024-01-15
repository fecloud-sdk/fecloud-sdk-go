package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TagsBody struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o TagsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsBody struct{}"
	}

	return strings.Join([]string{"TagsBody", string(data)}, " ")
}
