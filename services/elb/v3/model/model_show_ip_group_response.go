package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowIpGroupResponse Response Object
type ShowIpGroupResponse struct {
	Ipgroup *IpGroup `json:"ipgroup,omitempty"`

	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowIpGroupResponse", string(data)}, " ")
}
