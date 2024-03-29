package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateMembersRequestBody struct {
	Members []BatchCreateMembersOption `json:"members"`
}

func (o BatchCreateMembersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateMembersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateMembersRequestBody", string(data)}, " ")
}
