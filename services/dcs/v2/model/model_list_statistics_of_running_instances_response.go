package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListStatisticsOfRunningInstancesResponse struct {
	Statistics     *[]InstanceStatistic `json:"statistics,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListStatisticsOfRunningInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStatisticsOfRunningInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListStatisticsOfRunningInstancesResponse", string(data)}, " ")
}
