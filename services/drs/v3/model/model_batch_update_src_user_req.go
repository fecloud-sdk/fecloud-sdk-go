package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchUpdateSrcUserReq 批量更新迁移用户请求体
type BatchUpdateSrcUserReq struct {

	// 批量更新迁移用户请求列表
	Jobs []UpdateUserReq `json:"jobs"`
}

func (o BatchUpdateSrcUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateSrcUserReq struct{}"
	}

	return strings.Join([]string{"BatchUpdateSrcUserReq", string(data)}, " ")
}
