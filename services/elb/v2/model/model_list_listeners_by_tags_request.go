package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListListenersByTagsRequest struct {
	Body *ListListenersByTagsRequestBody `json:"body,omitempty"`
}

func (o ListListenersByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListListenersByTagsRequest", string(data)}, " ")
}
