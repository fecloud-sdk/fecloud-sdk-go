package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteServerTagsRequest struct {
	ServerId string `json:"server_id"`

	Body *BatchDeleteServerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteServerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerTagsRequest", string(data)}, " ")
}
