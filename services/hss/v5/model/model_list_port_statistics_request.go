package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPortStatisticsRequest struct {
	Port *int32 `json:"port,omitempty"`

	Type *string `json:"type,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListPortStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListPortStatisticsRequest", string(data)}, " ")
}
