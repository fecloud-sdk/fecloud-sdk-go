package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HostVulOperateInfo struct {
	HostId string `json:"host_id"`

	VulIdList []string `json:"vul_id_list"`
}

func (o HostVulOperateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostVulOperateInfo struct{}"
	}

	return strings.Join([]string{"HostVulOperateInfo", string(data)}, " ")
}
