package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TmsTag struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o TmsTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsTag struct{}"
	}

	return strings.Join([]string{"TmsTag", string(data)}, " ")
}
