package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowTransitIpTagsResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTransitIpTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransitIpTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowTransitIpTagsResponse", string(data)}, " ")
}
