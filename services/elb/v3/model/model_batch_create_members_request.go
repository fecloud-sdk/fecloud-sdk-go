package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchCreateMembersRequest Request Object
type BatchCreateMembersRequest struct {

	// 后端服务器组ID。
	PoolId string `json:"pool_id"`

	Body *BatchCreateMembersRequestBody `json:"body,omitempty"`
}

func (o BatchCreateMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateMembersRequest", string(data)}, " ")
}
