package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPrivateNatTagsResponse struct {
	RequestId *string `json:"request_id,omitempty"`

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
