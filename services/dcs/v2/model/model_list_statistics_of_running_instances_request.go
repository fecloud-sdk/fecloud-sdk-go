package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListStatisticsOfRunningInstancesRequest struct {
}

func (o ListStatisticsOfRunningInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStatisticsOfRunningInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListStatisticsOfRunningInstancesRequest", string(data)}, " ")
}
