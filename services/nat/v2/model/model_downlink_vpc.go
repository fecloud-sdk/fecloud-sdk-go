package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DownlinkVpc struct {
	VpcId string `json:"vpc_id"`

	VirsubnetId string `json:"virsubnet_id"`
}

func (o DownlinkVpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownlinkVpc struct{}"
	}

	return strings.Join([]string{"DownlinkVpc", string(data)}, " ")
}
