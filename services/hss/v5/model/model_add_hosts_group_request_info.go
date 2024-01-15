package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddHostsGroupRequestInfo struct {
	GroupName string `json:"group_name"`

	HostIdList []string `json:"host_id_list"`
}

func (o AddHostsGroupRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddHostsGroupRequestInfo struct{}"
	}

	return strings.Join([]string{"AddHostsGroupRequestInfo", string(data)}, " ")
}
