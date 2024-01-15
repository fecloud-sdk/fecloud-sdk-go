package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPrivateNatTagsResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Tags           *[]Tags `json:"tags,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPrivateNatTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateNatTagsResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateNatTagsResponse", string(data)}, " ")
}
