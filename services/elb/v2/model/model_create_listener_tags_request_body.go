package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateListenerTagsRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateListenerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateListenerTagsRequestBody", string(data)}, " ")
}
