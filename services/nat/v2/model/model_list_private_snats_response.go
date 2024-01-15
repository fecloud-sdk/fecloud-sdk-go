package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPrivateSnatsResponse struct {
	SnatRules *[]PrivateSnat `json:"snat_rules,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPrivateSnatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateSnatsResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateSnatsResponse", string(data)}, " ")
}
