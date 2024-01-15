package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePublicZoneReq struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	ZoneType *string `json:"zone_type,omitempty"`

	Email *string `json:"email,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreatePublicZoneReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicZoneReq struct{}"
	}

	return strings.Join([]string{"CreatePublicZoneReq", string(data)}, " ")
}
