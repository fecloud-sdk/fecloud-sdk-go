package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchUpdatePoliciesPriorityRequestBody This is a auto create Body Object
type BatchUpdatePoliciesPriorityRequestBody struct {

	// 转发策略的结构体
	L7policies *[]BatchUpdatePriorityRequestBody `json:"l7policies,omitempty"`
}

func (o BatchUpdatePoliciesPriorityRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePoliciesPriorityRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdatePoliciesPriorityRequestBody", string(data)}, " ")
}
