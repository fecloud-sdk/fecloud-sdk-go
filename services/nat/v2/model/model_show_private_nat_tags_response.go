package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowPrivateNatTagsResponse Response Object
type ShowPrivateNatTagsResponse struct {

	// 请求id。
	RequestId *string `json:"request_id,omitempty"`

	// 标签。
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowPrivateNatTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateNatTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivateNatTagsResponse", string(data)}, " ")
}
