package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// RestartInstanceResponse Response Object
type RestartInstanceResponse struct {

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestartInstanceResponse", string(data)}, " ")
}
