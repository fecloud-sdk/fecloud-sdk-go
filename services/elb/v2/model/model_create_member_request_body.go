package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateMemberRequestBody struct {
	Member *CreateMemberReq `json:"member"`
}

func (o CreateMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMemberRequestBody struct{}"
	}

	return strings.Join([]string{"CreateMemberRequestBody", string(data)}, " ")
}
