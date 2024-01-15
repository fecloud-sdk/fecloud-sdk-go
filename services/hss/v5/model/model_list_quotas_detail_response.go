package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListQuotasDetailResponse struct {
	PacketCycleNum *int32 `json:"packet_cycle_num,omitempty"`

	OnDemandNum *int32 `json:"on_demand_num,omitempty"`

	UsedNum *int32 `json:"used_num,omitempty"`

	IdleNum *int32 `json:"idle_num,omitempty"`

	NormalNum *int32 `json:"normal_num,omitempty"`

	ExpiredNum *int32 `json:"expired_num,omitempty"`

	FreezeNum *int32 `json:"freeze_num,omitempty"`

	QuotaStatisticsList *[]QuotaStatisticsResponseInfo `json:"quota_statistics_list,omitempty"`

	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]QuotaResourcesResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListQuotasDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotasDetailResponse struct{}"
	}

	return strings.Join([]string{"ListQuotasDetailResponse", string(data)}, " ")
}
