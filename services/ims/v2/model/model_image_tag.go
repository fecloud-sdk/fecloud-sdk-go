package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ImageTag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o ImageTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageTag struct{}"
	}

	return strings.Join([]string{"ImageTag", string(data)}, " ")
}
