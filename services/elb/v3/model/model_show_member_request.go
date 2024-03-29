package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowMemberRequest struct {
	PoolId string `json:"pool_id"`

	MemberId string `json:"member_id"`
}

func (o ShowMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMemberRequest struct{}"
	}

	return strings.Join([]string{"ShowMemberRequest", string(data)}, " ")
}
