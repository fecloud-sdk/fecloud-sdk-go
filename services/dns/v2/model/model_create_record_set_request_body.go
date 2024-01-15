package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateRecordSetRequestBody struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Type string `json:"type"`

	Status *string `json:"status,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`

	Records []string `json:"records"`

	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateRecordSetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRecordSetRequestBody", string(data)}, " ")
}
