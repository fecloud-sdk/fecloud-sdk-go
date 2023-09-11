package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// RevokeResponse Response Object
type RevokeResponse struct {

	// 操作结果。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RevokeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeResponse struct{}"
	}

	return strings.Join([]string{"RevokeResponse", string(data)}, " ")
}
