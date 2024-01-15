package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PolicyGroupResponseInfo struct {
	GroupName *string `json:"group_name,omitempty"`

	GroupId *string `json:"group_id,omitempty"`

	Description *string `json:"description,omitempty"`

	Deletable *bool `json:"deletable,omitempty"`

	HostNum *int32 `json:"host_num,omitempty"`

	DefaultGroup *bool `json:"default_group,omitempty"`

	SupportOs *string `json:"support_os,omitempty"`

	SupportVersion *string `json:"support_version,omitempty"`
}

func (o PolicyGroupResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroupResponseInfo struct{}"
	}

	return strings.Join([]string{"PolicyGroupResponseInfo", string(data)}, " ")
}
