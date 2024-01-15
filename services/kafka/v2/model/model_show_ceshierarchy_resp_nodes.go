package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCeshierarchyRespNodes struct {
	Name *string `json:"name,omitempty"`
}

func (o ShowCeshierarchyRespNodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespNodes struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespNodes", string(data)}, " ")
}
