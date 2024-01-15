package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePrivateZoneReq struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	ZoneType string `json:"zone_type"`

	Email *string `json:"email,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`

	Router *Router `json:"router"`

	ProxyPattern *string `json:"proxy_pattern,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreatePrivateZoneReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateZoneReq struct{}"
	}

	return strings.Join([]string{"CreatePrivateZoneReq", string(data)}, " ")
}
