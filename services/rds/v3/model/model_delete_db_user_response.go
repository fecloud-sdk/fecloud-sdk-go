package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteDbUserResponse Response Object
type DeleteDbUserResponse struct {

	// 操作结果。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDbUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteDbUserResponse", string(data)}, " ")
}
