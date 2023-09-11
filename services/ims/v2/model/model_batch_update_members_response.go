package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchUpdateMembersResponse Response Object
type BatchUpdateMembersResponse struct {

	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdateMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateMembersResponse", string(data)}, " ")
}
