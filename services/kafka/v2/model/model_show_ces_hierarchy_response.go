package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCesHierarchyResponse struct {
	Dimensions *[]ShowCeshierarchyRespDimensions `json:"dimensions,omitempty"`

	InstanceIds *[]ShowCeshierarchyRespInstanceIds `json:"instance_ids,omitempty"`

	Nodes *[]ShowCeshierarchyRespNodes `json:"nodes,omitempty"`

	Queues *[]ShowCeshierarchyRespQueues `json:"queues,omitempty"`

	Groups         *[]ShowCeshierarchyRespGroups `json:"groups,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowCesHierarchyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCesHierarchyResponse struct{}"
	}

	return strings.Join([]string{"ShowCesHierarchyResponse", string(data)}, " ")
}
