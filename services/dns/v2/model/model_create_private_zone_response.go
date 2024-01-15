package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePrivateZoneResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Email *string `json:"email,omitempty"`

	ZoneType *string `json:"zone_type,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`

	Serial *int32 `json:"serial,omitempty"`

	Status *string `json:"status,omitempty"`

	RecordNum *int32 `json:"record_num,omitempty"`

	ProxyPattern *string `json:"proxy_pattern,omitempty"`

	PoolId *string `json:"pool_id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	Masters *[]string `json:"masters,omitempty"`

	Router         *RouterWithStatus `json:"router,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreatePrivateZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateZoneResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivateZoneResponse", string(data)}, " ")
}
