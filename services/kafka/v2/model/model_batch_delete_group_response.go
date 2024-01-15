package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteGroupResponse struct {
	FailedGroups *[]BatchDeleteGroupRespFailedGroups `json:"failed_groups,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchDeleteGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteGroupResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteGroupResponse", string(data)}, " ")
}
