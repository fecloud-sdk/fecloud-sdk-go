package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PoolsInStatusResp struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Members []MembersInStatusResp `json:"members"`

	OperatingStatus string `json:"operating_status"`

	ProvisioningStatus string `json:"provisioning_status"`

	Healthmonitor *HealthmonitorsInStatusResp `json:"healthmonitor"`
}

func (o PoolsInStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolsInStatusResp struct{}"
	}

	return strings.Join([]string{"PoolsInStatusResp", string(data)}, " ")
}
