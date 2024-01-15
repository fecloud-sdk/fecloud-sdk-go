package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PrivateTag struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o PrivateTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateTag struct{}"
	}

	return strings.Join([]string{"PrivateTag", string(data)}, " ")
}
