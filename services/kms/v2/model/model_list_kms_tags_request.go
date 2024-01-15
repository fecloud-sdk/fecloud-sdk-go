package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListKmsTagsRequest struct {
}

func (o ListKmsTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKmsTagsRequest struct{}"
	}

	return strings.Join([]string{"ListKmsTagsRequest", string(data)}, " ")
}
