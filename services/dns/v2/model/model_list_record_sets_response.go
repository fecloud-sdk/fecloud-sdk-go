package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRecordSetsResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Recordsets *[]ListRecordSetsWithTags `json:"recordsets,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRecordSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsResponse struct{}"
	}

	return strings.Join([]string{"ListRecordSetsResponse", string(data)}, " ")
}
