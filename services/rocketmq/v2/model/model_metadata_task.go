package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MetadataTask struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	StartDate *string `json:"start_date,omitempty"`

	Status *string `json:"status,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o MetadataTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetadataTask struct{}"
	}

	return strings.Join([]string{"MetadataTask", string(data)}, " ")
}
