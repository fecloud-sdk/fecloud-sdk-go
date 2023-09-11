package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// RestoreExistInstanceResponse Response Object
type RestoreExistInstanceResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreExistInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreExistInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestoreExistInstanceResponse", string(data)}, " ")
}
