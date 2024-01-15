package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ConsumerList struct {
	Topics *[]string `json:"topics,omitempty"`

	Total *int32 `json:"total,omitempty"`
}

func (o ConsumerList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerList struct{}"
	}

	return strings.Join([]string{"ConsumerList", string(data)}, " ")
}
