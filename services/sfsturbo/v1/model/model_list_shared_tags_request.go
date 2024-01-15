package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSharedTagsRequest struct {
	ContentType string `json:"Content-Type"`
}

func (o ListSharedTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharedTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSharedTagsRequest", string(data)}, " ")
}
