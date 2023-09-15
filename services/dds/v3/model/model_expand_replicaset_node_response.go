package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ExpandReplicasetNodeResponse Response Object
type ExpandReplicasetNodeResponse struct {

	// 任务ID，仅按需实例返回该参数。
	JobId *string `json:"job_id,omitempty"`

	// 订单ID，仅包周期实例返回该参数。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandReplicasetNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandReplicasetNodeResponse struct{}"
	}

	return strings.Join([]string{"ExpandReplicasetNodeResponse", string(data)}, " ")
}
