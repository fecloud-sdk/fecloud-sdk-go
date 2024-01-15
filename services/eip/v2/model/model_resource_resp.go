package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResourceResp struct {

	// 资源配额对象
	Resources []QuotaShowResp `json:"resources"`
}

func (o ResourceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceResp struct{}"
	}

	return strings.Join([]string{"ResourceResp", string(data)}, " ")
}
