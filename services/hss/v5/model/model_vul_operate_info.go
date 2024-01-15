package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type VulOperateInfo struct {
	VulId string `json:"vul_id"`

	HostIdList []string `json:"host_id_list"`
}

func (o VulOperateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulOperateInfo struct{}"
	}

	return strings.Join([]string{"VulOperateInfo", string(data)}, " ")
}
