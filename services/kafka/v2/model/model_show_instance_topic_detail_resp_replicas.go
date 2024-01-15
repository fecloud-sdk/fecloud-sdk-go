package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowInstanceTopicDetailRespReplicas struct {
	Broker *int32 `json:"broker,omitempty"`

	Leader *bool `json:"leader,omitempty"`

	InSync *bool `json:"in_sync,omitempty"`

	Size *int32 `json:"size,omitempty"`

	Lag *int32 `json:"lag,omitempty"`
}

func (o ShowInstanceTopicDetailRespReplicas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTopicDetailRespReplicas struct{}"
	}

	return strings.Join([]string{"ShowInstanceTopicDetailRespReplicas", string(data)}, " ")
}
