package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateMembersRequestBody struct {
	Members []BatchUpdateMembersOption `json:"members"`
}

func (o BatchUpdateMembersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateMembersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateMembersRequestBody", string(data)}, " ")
}
