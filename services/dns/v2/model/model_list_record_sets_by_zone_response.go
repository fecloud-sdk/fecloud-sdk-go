package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRecordSetsByZoneResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Recordsets *[]ListRecordSets `json:"recordsets,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRecordSetsByZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsByZoneResponse struct{}"
	}

	return strings.Join([]string{"ListRecordSetsByZoneResponse", string(data)}, " ")
}
