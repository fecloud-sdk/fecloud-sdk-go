package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QuotaStatisticsResponseInfo struct {
	Version *string `json:"version,omitempty"`

	TotalNum *int32 `json:"total_num,omitempty"`
}

func (o QuotaStatisticsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaStatisticsResponseInfo struct{}"
	}

	return strings.Join([]string{"QuotaStatisticsResponseInfo", string(data)}, " ")
}
