package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// PoolRef
type PoolRef struct {

	// 后端服务器组ID。
	Id string `json:"id"`
}

func (o PoolRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolRef struct{}"
	}

	return strings.Join([]string{"PoolRef", string(data)}, " ")
}