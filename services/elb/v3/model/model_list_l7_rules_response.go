package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListL7RulesResponse Response Object
type ListL7RulesResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 规则对象列表。
	Rules          *[]L7Rule `json:"rules,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListL7RulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7RulesResponse struct{}"
	}

	return strings.Join([]string{"ListL7RulesResponse", string(data)}, " ")
}