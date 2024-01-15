package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListListenerTagsResponse struct {
	Tags           *[]ListTag `json:"tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListListenerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"ListListenerTagsResponse", string(data)}, " ")
}
