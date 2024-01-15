package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EventId struct {
}

func (o EventId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventId struct{}"
	}

	return strings.Join([]string{"EventId", string(data)}, " ")
}
