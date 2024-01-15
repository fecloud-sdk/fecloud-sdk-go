package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ErrorResponseBody struct {
	ErrorCode string `json:"error_code"`

	ErrorMsg string `json:"error_msg"`
}

func (o ErrorResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorResponseBody struct{}"
	}

	return strings.Join([]string{"ErrorResponseBody", string(data)}, " ")
}
