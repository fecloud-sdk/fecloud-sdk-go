package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DownlinkVpcOption struct {
	VirsubnetId string `json:"virsubnet_id"`
}

func (o DownlinkVpcOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownlinkVpcOption struct{}"
	}

	return strings.Join([]string{"DownlinkVpcOption", string(data)}, " ")
}
