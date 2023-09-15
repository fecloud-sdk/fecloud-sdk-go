package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchDeleteListenerTagsResponse Response Object
type BatchDeleteListenerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteListenerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteListenerTagsResponse", string(data)}, " ")
}
