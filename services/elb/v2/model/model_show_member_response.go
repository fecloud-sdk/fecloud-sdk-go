package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowMemberResponse Response Object
type ShowMemberResponse struct {
	Member         *MemberResp `json:"member,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMemberResponse struct{}"
	}

	return strings.Join([]string{"ShowMemberResponse", string(data)}, " ")
}
