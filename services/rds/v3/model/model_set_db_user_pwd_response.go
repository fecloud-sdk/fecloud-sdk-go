package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// SetDbUserPwdResponse Response Object
type SetDbUserPwdResponse struct {

	// 操作结果。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetDbUserPwdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDbUserPwdResponse struct{}"
	}

	return strings.Join([]string{"SetDbUserPwdResponse", string(data)}, " ")
}
