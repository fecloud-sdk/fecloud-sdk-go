package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdatePortResponse Response Object
type UpdatePortResponse struct {

	// 任务ID
	WorkflowId     *string `json:"workflowId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortResponse struct{}"
	}

	return strings.Join([]string{"UpdatePortResponse", string(data)}, " ")
}
