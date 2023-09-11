package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateDbUserResponse Response Object
type CreateDbUserResponse struct {

	// 操作结果。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDbUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbUserResponse struct{}"
	}

	return strings.Join([]string{"CreateDbUserResponse", string(data)}, " ")
}
