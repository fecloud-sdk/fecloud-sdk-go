package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListMembersResponse struct {
	Members        *[]MemberResp `json:"members,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMembersResponse struct{}"
	}

	return strings.Join([]string{"ListMembersResponse", string(data)}, " ")
}
