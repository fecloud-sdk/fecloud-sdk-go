package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateKmsTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateKmsTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateKmsTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateKmsTagsResponse", string(data)}, " ")
}
