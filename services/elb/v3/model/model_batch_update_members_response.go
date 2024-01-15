package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateMembersResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Members        *[]BatchUpdateMember `json:"members,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o BatchUpdateMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateMembersResponse", string(data)}, " ")
}
