package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteServerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteServerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerTagsResponse", string(data)}, " ")
}
