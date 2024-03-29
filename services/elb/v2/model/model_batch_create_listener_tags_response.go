package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateListenerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateListenerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateListenerTagsResponse", string(data)}, " ")
}
