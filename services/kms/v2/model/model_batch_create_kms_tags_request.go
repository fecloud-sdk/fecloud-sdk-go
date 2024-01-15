package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateKmsTagsRequest struct {
	KeyId string `json:"key_id"`

	Body *BatchCreateKmsTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateKmsTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateKmsTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateKmsTagsRequest", string(data)}, " ")
}
