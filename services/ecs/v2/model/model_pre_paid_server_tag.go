package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PrePaidServerTag struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o PrePaidServerTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServerTag struct{}"
	}

	return strings.Join([]string{"PrePaidServerTag", string(data)}, " ")
}
