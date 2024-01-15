package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAllTagsRequest struct {
}

func (o ListAllTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTagsRequest struct{}"
	}

	return strings.Join([]string{"ListAllTagsRequest", string(data)}, " ")
}
