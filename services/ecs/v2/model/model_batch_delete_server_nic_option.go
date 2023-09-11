package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchDeleteServerNicOption
type BatchDeleteServerNicOption struct {

	// 网卡Port ID。
	Id string `json:"id"`
}

func (o BatchDeleteServerNicOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerNicOption struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerNicOption", string(data)}, " ")
}
