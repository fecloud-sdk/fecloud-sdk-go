package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateShareRequestBody struct {
	Share *Share `json:"share"`
}

func (o CreateShareRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareRequestBody struct{}"
	}

	return strings.Join([]string{"CreateShareRequestBody", string(data)}, " ")
}
