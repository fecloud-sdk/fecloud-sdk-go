package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLogGroupsRequest struct {
	ContentType string `json:"Content-Type"`
}

func (o ListLogGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListLogGroupsRequest", string(data)}, " ")
}
