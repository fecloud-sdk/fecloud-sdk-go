package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateMemberResponse Response Object
type CreateMemberResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	Member         *Member `json:"member,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMemberResponse struct{}"
	}

	return strings.Join([]string{"CreateMemberResponse", string(data)}, " ")
}