package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowPoolRequest Request Object
type ShowPoolRequest struct {

	// 后端服务器组ID。
	PoolId string `json:"pool_id"`
}

func (o ShowPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolRequest struct{}"
	}

	return strings.Join([]string{"ShowPoolRequest", string(data)}, " ")
}
