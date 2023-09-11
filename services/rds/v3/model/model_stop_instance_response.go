package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// StopInstanceResponse Response Object
type StopInstanceResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopInstanceResponse struct{}"
	}

	return strings.Join([]string{"StopInstanceResponse", string(data)}, " ")
}
