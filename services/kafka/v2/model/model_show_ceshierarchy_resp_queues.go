package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCeshierarchyRespQueues struct {
	Name *string `json:"name,omitempty"`

	Partitions *[]ShowCeshierarchyRespPartitions `json:"partitions,omitempty"`
}

func (o ShowCeshierarchyRespQueues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespQueues struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespQueues", string(data)}, " ")
}
