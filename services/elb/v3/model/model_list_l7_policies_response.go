package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListL7PoliciesResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	L7policies     *[]L7Policy `json:"l7policies,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListL7PoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7PoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListL7PoliciesResponse", string(data)}, " ")
}
