package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPtrRecordsResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Floatingips    *[]ListPtrRecordsFloatingResp `json:"floatingips,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListPtrRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPtrRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListPtrRecordsResponse", string(data)}, " ")
}
