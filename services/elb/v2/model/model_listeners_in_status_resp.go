package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListenersInStatusResp struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Pools []PoolsInStatusResp `json:"pools"`

	L7policies []L7policiesInStatusResp `json:"l7policies"`

	OperatingStatus string `json:"operating_status"`

	ProvisioningStatus string `json:"provisioning_status"`
}

func (o ListenersInStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenersInStatusResp struct{}"
	}

	return strings.Join([]string{"ListenersInStatusResp", string(data)}, " ")
}
