package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResourceTag struct {

	// 标签key值。
	Key string `json:"key"`

	// 标签value。
	Value string `json:"value"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
