package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateL7PolicyResponse Response Object
type CreateL7PolicyResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	L7policy       *L7Policy `json:"l7policy,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateL7PolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7PolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateL7PolicyResponse", string(data)}, " ")
}
