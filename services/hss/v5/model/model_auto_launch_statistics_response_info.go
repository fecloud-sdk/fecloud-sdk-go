package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AutoLaunchStatisticsResponseInfo struct {
	Name *string `json:"name,omitempty"`

	Type *string `json:"type,omitempty"`

	Num *int32 `json:"num,omitempty"`
}

func (o AutoLaunchStatisticsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoLaunchStatisticsResponseInfo struct{}"
	}

	return strings.Join([]string{"AutoLaunchStatisticsResponseInfo", string(data)}, " ")
}
