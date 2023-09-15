package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// PreoccupyIp
type PreoccupyIp struct {

	// 预占IP总数
	Total int32 `json:"total"`
}

func (o PreoccupyIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreoccupyIp struct{}"
	}

	return strings.Join([]string{"PreoccupyIp", string(data)}, " ")
}
