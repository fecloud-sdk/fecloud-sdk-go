package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCeshierarchyRespInstanceIds struct {
	Name *string `json:"name,omitempty"`
}

func (o ShowCeshierarchyRespInstanceIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespInstanceIds struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespInstanceIds", string(data)}, " ")
}
