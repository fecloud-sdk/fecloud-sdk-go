package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRecordSetsWithLineResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Recordsets *[]QueryRecordSetWithLineAndTagsResp `json:"recordsets,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRecordSetsWithLineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsWithLineResponse struct{}"
	}

	return strings.Join([]string{"ListRecordSetsWithLineResponse", string(data)}, " ")
}
