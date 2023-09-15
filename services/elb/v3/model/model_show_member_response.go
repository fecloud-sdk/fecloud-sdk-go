package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowMemberResponse Response Object
type ShowMemberResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	Member         *Member `json:"member,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMemberResponse struct{}"
	}

	return strings.Join([]string{"ShowMemberResponse", string(data)}, " ")
}
