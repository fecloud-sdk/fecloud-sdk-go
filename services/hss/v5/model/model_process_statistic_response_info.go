package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ProcessStatisticResponseInfo struct {
	Path *string `json:"path,omitempty"`

	Num *int32 `json:"num,omitempty"`
}

func (o ProcessStatisticResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessStatisticResponseInfo struct{}"
	}

	return strings.Join([]string{"ProcessStatisticResponseInfo", string(data)}, " ")
}
