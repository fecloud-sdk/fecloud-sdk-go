package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeletePoolRequest Request Object
type DeletePoolRequest struct {

	// 后端服务器组ID。
	PoolId string `json:"pool_id"`
}

func (o DeletePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePoolRequest struct{}"
	}

	return strings.Join([]string{"DeletePoolRequest", string(data)}, " ")
}
