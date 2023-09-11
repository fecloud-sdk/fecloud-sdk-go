package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateImageResponse Response Object
type CreateImageResponse struct {

	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageResponse struct{}"
	}

	return strings.Join([]string{"CreateImageResponse", string(data)}, " ")
}