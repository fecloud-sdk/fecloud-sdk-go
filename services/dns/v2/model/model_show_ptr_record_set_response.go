package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPtrRecordSetResponse struct {
	Id *string `json:"id,omitempty"`

	Ptrdname *string `json:"ptrdname,omitempty"`

	Description *string `json:"description,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`

	Address *string `json:"address,omitempty"`

	Status *string `json:"status,omitempty"`

	Action *string `json:"action,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowPtrRecordSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPtrRecordSetResponse struct{}"
	}

	return strings.Join([]string{"ShowPtrRecordSetResponse", string(data)}, " ")
}
