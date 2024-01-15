package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UserStatisticInfoResponseInfo struct {
	UserName *string `json:"user_name,omitempty"`

	Num *int32 `json:"num,omitempty"`
}

func (o UserStatisticInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserStatisticInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"UserStatisticInfoResponseInfo", string(data)}, " ")
}
