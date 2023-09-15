package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdatePoolRequest Request Object
type UpdatePoolRequest struct {

	// 后端服务器组ID。
	PoolId string `json:"pool_id"`

	Body *UpdatePoolRequestBody `json:"body,omitempty"`
}

func (o UpdatePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolRequest struct{}"
	}

	return strings.Join([]string{"UpdatePoolRequest", string(data)}, " ")
}