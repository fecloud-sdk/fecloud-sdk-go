package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPublicZonesRequest struct {
	Type *string `json:"type,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Tags *string `json:"tags,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	SearchMode *string `json:"search_mode,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListPublicZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicZonesRequest struct{}"
	}

	return strings.Join([]string{"ListPublicZonesRequest", string(data)}, " ")
}
