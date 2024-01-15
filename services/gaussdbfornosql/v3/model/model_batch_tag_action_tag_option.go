package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchTagActionTagOption struct {
	Key string `json:"key"`

	Value *string `json:"value,omitempty"`
}

func (o BatchTagActionTagOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagActionTagOption struct{}"
	}

	return strings.Join([]string{"BatchTagActionTagOption", string(data)}, " ")
}
