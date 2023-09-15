package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteSessionRequestBody struct {

	// 需要终结的会话ID列表。
	Sessions []string `json:"sessions"`
}

func (o DeleteSessionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSessionRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteSessionRequestBody", string(data)}, " ")
}
