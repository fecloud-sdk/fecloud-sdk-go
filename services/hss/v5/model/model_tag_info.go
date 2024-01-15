package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TagInfo struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o TagInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagInfo struct{}"
	}

	return strings.Join([]string{"TagInfo", string(data)}, " ")
}
