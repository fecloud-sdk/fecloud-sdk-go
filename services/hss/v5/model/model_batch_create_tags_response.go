package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateTagsResponse", string(data)}, " ")
}
