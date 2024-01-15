package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCeshierarchyRespPartitions struct {
	Name *string `json:"name,omitempty"`
}

func (o ShowCeshierarchyRespPartitions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespPartitions struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespPartitions", string(data)}, " ")
}
