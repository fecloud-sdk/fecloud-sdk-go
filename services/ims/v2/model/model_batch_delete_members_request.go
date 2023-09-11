package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchDeleteMembersRequest Request Object
type BatchDeleteMembersRequest struct {
	Body *BatchAddMembersRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersRequest", string(data)}, " ")
}
