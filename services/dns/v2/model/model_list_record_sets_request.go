package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRecordSetsRequest struct {
	ZoneType *string `json:"zone_type,omitempty"`

	Marker *string `json:"marker,omitempty"`

	SearchMode *string `json:"search_mode,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Tags *string `json:"tags,omitempty"`

	Status *string `json:"status,omitempty"`

	Type *string `json:"type,omitempty"`

	Name *string `json:"name,omitempty"`

	Id *string `json:"id,omitempty"`

	Records *string `json:"records,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListRecordSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsRequest struct{}"
	}

	return strings.Join([]string{"ListRecordSetsRequest", string(data)}, " ")
}
