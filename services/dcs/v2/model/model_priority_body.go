package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PriorityBody struct {
	SlavePriorityWeight int32 `json:"slave_priority_weight"`
}

func (o PriorityBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PriorityBody struct{}"
	}

	return strings.Join([]string{"PriorityBody", string(data)}, " ")
}
