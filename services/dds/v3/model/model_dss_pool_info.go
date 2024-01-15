package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DssPoolInfo struct {
	AzName string `json:"az_name"`

	FreeCapacityGb string `json:"free_capacity_gb"`

	DssPoolVolumeType string `json:"dss_pool_volume_type"`

	DssPoolId string `json:"dss_pool_id"`

	DssPoolStatus string `json:"dss_pool_status"`
}

func (o DssPoolInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DssPoolInfo struct{}"
	}

	return strings.Join([]string{"DssPoolInfo", string(data)}, " ")
}
