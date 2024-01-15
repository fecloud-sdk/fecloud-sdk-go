package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCeshierarchyRespQueues1 struct {
	Name *string `json:"name,omitempty"`

	Partitions *[]ShowCeshierarchyRespPartitions `json:"partitions,omitempty"`
}

func (o ShowCeshierarchyRespQueues1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespQueues1 struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespQueues1", string(data)}, " ")
}
