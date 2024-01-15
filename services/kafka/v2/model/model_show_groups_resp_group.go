package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowGroupsRespGroup struct {
	GroupId *string `json:"group_id,omitempty"`

	State *string `json:"state,omitempty"`

	CoordinatorId *int32 `json:"coordinator_id,omitempty"`

	Members *[]ShowGroupsRespGroupMembers `json:"members,omitempty"`

	GroupMessageOffsets *[]ShowGroupsRespGroupGroupMessageOffsets `json:"group_message_offsets,omitempty"`

	AssignmentStrategy *string `json:"assignment_strategy,omitempty"`
}

func (o ShowGroupsRespGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsRespGroup struct{}"
	}

	return strings.Join([]string{"ShowGroupsRespGroup", string(data)}, " ")
}
