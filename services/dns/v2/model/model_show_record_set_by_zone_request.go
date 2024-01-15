package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowRecordSetByZoneRequest struct {
	ZoneId string `json:"zone_id"`

	Marker *string `json:"marker,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	LineId *string `json:"line_id,omitempty"`

	Tags *string `json:"tags,omitempty"`

	Status *string `json:"status,omitempty"`

	Type *string `json:"type,omitempty"`

	Name *string `json:"name,omitempty"`

	Id *string `json:"id,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	SortDir *string `json:"sort_dir,omitempty"`

	SearchMode *string `json:"search_mode,omitempty"`
}

func (o ShowRecordSetByZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordSetByZoneRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordSetByZoneRequest", string(data)}, " ")
}
