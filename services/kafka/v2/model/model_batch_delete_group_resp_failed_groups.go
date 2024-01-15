package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteGroupRespFailedGroups struct {
	GroupId *string `json:"group_id,omitempty"`

	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o BatchDeleteGroupRespFailedGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteGroupRespFailedGroups struct{}"
	}

	return strings.Join([]string{"BatchDeleteGroupRespFailedGroups", string(data)}, " ")
}
