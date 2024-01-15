package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResourceTags struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o ResourceTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTags struct{}"
	}

	return strings.Join([]string{"ResourceTags", string(data)}, " ")
}
