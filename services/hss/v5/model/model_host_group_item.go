package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HostGroupItem struct {
	GroupId *string `json:"group_id,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	HostNum *int32 `json:"host_num,omitempty"`

	RiskHostNum *int32 `json:"risk_host_num,omitempty"`

	UnprotectHostNum *int32 `json:"unprotect_host_num,omitempty"`

	HostIdList *[]string `json:"host_id_list,omitempty"`

	IsOutside *bool `json:"is_outside,omitempty"`
}

func (o HostGroupItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostGroupItem struct{}"
	}

	return strings.Join([]string{"HostGroupItem", string(data)}, " ")
}
