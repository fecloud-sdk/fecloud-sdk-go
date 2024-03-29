package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type InstanceStatistics struct {
	Unhealth *int32 `json:"unhealth,omitempty"`

	Total *int32 `json:"total,omitempty"`

	TypeStatistics *int32 `json:"type_statistics,omitempty"`
}

func (o InstanceStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceStatistics struct{}"
	}

	return strings.Join([]string{"InstanceStatistics", string(data)}, " ")
}
