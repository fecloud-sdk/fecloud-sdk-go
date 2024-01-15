package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type InstanceTagResult struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o InstanceTagResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceTagResult struct{}"
	}

	return strings.Join([]string{"InstanceTagResult", string(data)}, " ")
}
