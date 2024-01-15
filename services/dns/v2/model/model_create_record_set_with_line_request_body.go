package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateRecordSetWithLineRequestBody struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Type string `json:"type"`

	Status *string `json:"status,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`

	Records *[]string `json:"records,omitempty"`

	Line *string `json:"line,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`

	Weight *int32 `json:"weight,omitempty"`

	AliasTarget *AliasTarget `json:"alias_target,omitempty"`
}

func (o CreateRecordSetWithLineRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetWithLineRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRecordSetWithLineRequestBody", string(data)}, " ")
}
