package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// StartResizeFlavorActionResponse Response Object
type StartResizeFlavorActionResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartResizeFlavorActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartResizeFlavorActionResponse struct{}"
	}

	return strings.Join([]string{"StartResizeFlavorActionResponse", string(data)}, " ")
}
