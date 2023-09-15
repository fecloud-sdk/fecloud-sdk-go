package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// SetBalancerWindowResponse Response Object
type SetBalancerWindowResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetBalancerWindowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBalancerWindowResponse struct{}"
	}

	return strings.Join([]string{"SetBalancerWindowResponse", string(data)}, " ")
}
