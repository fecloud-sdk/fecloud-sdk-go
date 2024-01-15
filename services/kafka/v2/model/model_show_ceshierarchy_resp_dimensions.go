package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCeshierarchyRespDimensions struct {
	Name *string `json:"name,omitempty"`

	Metrics *[]string `json:"metrics,omitempty"`

	KeyName *[]string `json:"key_name,omitempty"`

	DimRouter *[]string `json:"dim_router,omitempty"`

	Children *[]ShowCeshierarchyRespChildren `json:"children,omitempty"`
}

func (o ShowCeshierarchyRespDimensions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespDimensions struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespDimensions", string(data)}, " ")
}
