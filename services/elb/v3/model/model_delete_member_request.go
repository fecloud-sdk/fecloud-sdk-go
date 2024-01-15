package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteMemberRequest struct {
	PoolId string `json:"pool_id"`

	MemberId string `json:"member_id"`
}

func (o DeleteMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMemberRequest struct{}"
	}

	return strings.Join([]string{"DeleteMemberRequest", string(data)}, " ")
}
