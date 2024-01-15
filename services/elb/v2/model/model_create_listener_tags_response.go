package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateListenerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateListenerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateListenerTagsResponse", string(data)}, " ")
}
