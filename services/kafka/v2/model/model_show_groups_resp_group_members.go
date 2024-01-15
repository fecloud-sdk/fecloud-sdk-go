package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowGroupsRespGroupMembers struct {
	Host *string `json:"host,omitempty"`

	Assignment *[]ShowGroupsRespGroupAssignment `json:"assignment,omitempty"`

	MemberId *string `json:"member_id,omitempty"`

	ClientId *string `json:"client_id,omitempty"`
}

func (o ShowGroupsRespGroupMembers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsRespGroupMembers struct{}"
	}

	return strings.Join([]string{"ShowGroupsRespGroupMembers", string(data)}, " ")
}
