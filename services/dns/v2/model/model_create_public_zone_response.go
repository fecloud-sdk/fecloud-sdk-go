package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePublicZoneResponse struct {
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

	ProjectId *string `json:"project_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Masters        *[]string `json:"masters,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreatePublicZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicZoneResponse struct{}"
	}

	return strings.Join([]string{"CreatePublicZoneResponse", string(data)}, " ")
}
