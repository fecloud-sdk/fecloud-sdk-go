package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchDeleteIpListResponse Response Object
type BatchDeleteIpListResponse struct {
	Ipgroup *IpGroup `json:"ipgroup,omitempty"`

	// 请求ID。 注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteIpListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIpListResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteIpListResponse", string(data)}, " ")
}
