package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// SwitchoverReplicaSetResponse Response Object
type SwitchoverReplicaSetResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchoverReplicaSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchoverReplicaSetResponse struct{}"
	}

	return strings.Join([]string{"SwitchoverReplicaSetResponse", string(data)}, " ")
}
