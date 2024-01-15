package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPrivateDnatsResponse struct {
	DnatRules *[]PrivateDnat `json:"dnat_rules,omitempty"`

	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPrivateDnatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateDnatsResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateDnatsResponse", string(data)}, " ")
}
