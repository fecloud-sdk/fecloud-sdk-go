package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowListenerTagsResponse struct {
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowListenerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowListenerTagsResponse", string(data)}, " ")
}
