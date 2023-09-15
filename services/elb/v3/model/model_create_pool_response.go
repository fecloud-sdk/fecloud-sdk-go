package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreatePoolResponse Response Object
type CreatePoolResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	Pool           *Pool `json:"pool,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreatePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolResponse struct{}"
	}

	return strings.Join([]string{"CreatePoolResponse", string(data)}, " ")
}
