package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCeshierarchyRespGroups struct {
	Name *string `json:"name,omitempty"`

	Queues *[]ShowCeshierarchyRespQueues1 `json:"queues,omitempty"`
}

func (o ShowCeshierarchyRespGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespGroups struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespGroups", string(data)}, " ")
}
