package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCeshierarchyRespChildren struct {
	Name *string `json:"name,omitempty"`

	Metrics *[]string `json:"metrics,omitempty"`

	KeyName *[]string `json:"key_name,omitempty"`

	DimRouter *[]string `json:"dim_router,omitempty"`
}

func (o ShowCeshierarchyRespChildren) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespChildren struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespChildren", string(data)}, " ")
}
