package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListInstancesNodeResult struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Status string `json:"status"`

	Role string `json:"role"`

	SubnetId string `json:"subnet_id"`

	PrivateIp string `json:"private_ip"`

	PublicIp string `json:"public_ip"`

	SpecCode string `json:"spec_code"`

	AvailabilityZone string `json:"availability_zone"`

	SupportReduce bool `json:"support_reduce"`
}

func (o ListInstancesNodeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesNodeResult struct{}"
	}

	return strings.Join([]string{"ListInstancesNodeResult", string(data)}, " ")
}
