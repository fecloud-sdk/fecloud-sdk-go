package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type GroupInfoSimple struct {
	CreatedAt *int32 `json:"createdAt,omitempty"`

	GroupId *string `json:"group_id,omitempty"`

	State *string `json:"state,omitempty"`

	CoordinatorId *int32 `json:"coordinator_id,omitempty"`

	GroupDesc *string `json:"group_desc,omitempty"`

	Lag *int32 `json:"lag,omitempty"`
}

func (o GroupInfoSimple) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupInfoSimple struct{}"
	}

	return strings.Join([]string{"GroupInfoSimple", string(data)}, " ")
}
