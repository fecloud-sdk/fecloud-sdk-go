package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSinkTaskDetailRespTopicsInfo struct {
	Topic *string `json:"topic,omitempty"`

	Partitions *[]ShowSinkTaskDetailRespPartitions `json:"partitions,omitempty"`
}

func (o ShowSinkTaskDetailRespTopicsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSinkTaskDetailRespTopicsInfo struct{}"
	}

	return strings.Join([]string{"ShowSinkTaskDetailRespTopicsInfo", string(data)}, " ")
}
