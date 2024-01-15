package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateMemberResponse struct {
	Member         *MemberResp `json:"member,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMemberResponse struct{}"
	}

	return strings.Join([]string{"CreateMemberResponse", string(data)}, " ")
}
