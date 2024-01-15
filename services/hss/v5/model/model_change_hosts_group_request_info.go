package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeHostsGroupRequestInfo struct {
	GroupName *string `json:"group_name,omitempty"`

	GroupId string `json:"group_id"`

	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o ChangeHostsGroupRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeHostsGroupRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeHostsGroupRequestInfo", string(data)}, " ")
}
