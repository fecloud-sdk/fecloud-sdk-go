package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResourceTagInfo struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o ResourceTagInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTagInfo struct{}"
	}

	return strings.Join([]string{"ResourceTagInfo", string(data)}, " ")
}
