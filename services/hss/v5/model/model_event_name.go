package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EventName struct {
}

func (o EventName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventName struct{}"
	}

	return strings.Join([]string{"EventName", string(data)}, " ")
}
