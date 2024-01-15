package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EventClassId struct {
}

func (o EventClassId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventClassId struct{}"
	}

	return strings.Join([]string{"EventClassId", string(data)}, " ")
}
