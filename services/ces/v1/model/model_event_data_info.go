package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EventDataInfo struct {
	Type string `json:"type"`

	Timestamp int64 `json:"timestamp"`

	Value string `json:"value"`
}

func (o EventDataInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventDataInfo struct{}"
	}

	return strings.Join([]string{"EventDataInfo", string(data)}, " ")
}
