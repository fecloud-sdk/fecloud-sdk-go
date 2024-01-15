package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListListenerTagsRequest struct {
}

func (o ListListenerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"ListListenerTagsRequest", string(data)}, " ")
}
