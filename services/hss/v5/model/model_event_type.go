package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EventType struct {
}

func (o EventType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventType struct{}"
	}

	return strings.Join([]string{"EventType", string(data)}, " ")
}
