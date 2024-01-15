package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteRecordSetsResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	ZoneId *string `json:"zone_id,omitempty"`

	ZoneName *string `json:"zone_name,omitempty"`

	Type *string `json:"type,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`

	Records *[]string `json:"records,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	Status *string `json:"status,omitempty"`

	Default *bool `json:"default,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	Line *string `json:"line,omitempty"`

	Weight *int32 `json:"weight,omitempty"`

	HealthCheckId *string `json:"health_check_id,omitempty"`

	AliasTarget    *AliasTarget `json:"alias_target,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteRecordSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecordSetsResponse struct{}"
	}

	return strings.Join([]string{"DeleteRecordSetsResponse", string(data)}, " ")
}
