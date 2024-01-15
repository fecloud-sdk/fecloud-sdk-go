package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPublicZoneResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Email *string `json:"email,omitempty"`

	ZoneType *string `json:"zone_type,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`

	Serial *int32 `json:"serial,omitempty"`

	Status *string `json:"status,omitempty"`

	RecordNum *int32 `json:"record_num,omitempty"`

	PoolId *string `json:"pool_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Masters *[]string `json:"masters,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	ProjectId      *string `json:"project_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPublicZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicZoneResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicZoneResponse", string(data)}, " ")
}
