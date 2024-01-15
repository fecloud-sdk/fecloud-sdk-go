package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateSharedBandwidhRequestBody 创建共享带宽请求体
type CreateSharedBandwidhRequestBody struct {
	Bandwidth *CreateSharedBandwidthOption `json:"bandwidth"`
}

func (o CreateSharedBandwidhRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSharedBandwidhRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSharedBandwidhRequestBody", string(data)}, " ")
}
