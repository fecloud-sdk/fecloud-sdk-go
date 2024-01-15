package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeletePrivateZoneResponse struct {
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

	Masters *[]string `json:"masters,omitempty"`

	Routers        *[]RouterWithStatus `json:"routers,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o DeletePrivateZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateZoneResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateZoneResponse", string(data)}, " ")
}
