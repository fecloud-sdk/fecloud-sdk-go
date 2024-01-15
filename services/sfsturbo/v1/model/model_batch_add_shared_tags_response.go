package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchAddSharedTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddSharedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddSharedTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddSharedTagsResponse", string(data)}, " ")
}
