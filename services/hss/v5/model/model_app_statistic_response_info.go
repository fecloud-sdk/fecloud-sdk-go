package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AppStatisticResponseInfo struct {
	AppName *string `json:"app_name,omitempty"`

	Num *int32 `json:"num,omitempty"`
}

func (o AppStatisticResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppStatisticResponseInfo struct{}"
	}

	return strings.Join([]string{"AppStatisticResponseInfo", string(data)}, " ")
}
