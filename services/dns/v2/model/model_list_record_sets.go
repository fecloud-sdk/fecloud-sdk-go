package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRecordSets struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	ZoneId *string `json:"zone_id,omitempty"`

	ZoneName *string `json:"zone_name,omitempty"`

	Type *string `json:"type,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`

	Records *[]string `json:"records,omitempty"`

	CreateAt *string `json:"create_at,omitempty"`

	UpdateAt *string `json:"update_at,omitempty"`

	Status *string `json:"status,omitempty"`

	Default *bool `json:"default,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Links *PageLink `json:"links,omitempty"`
}

func (o ListRecordSets) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSets struct{}"
	}

	return strings.Join([]string{"ListRecordSets", string(data)}, " ")
}
