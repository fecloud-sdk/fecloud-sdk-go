package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListFlavorsRequest Request Object
type ListFlavorsRequest struct {

	// 可用区，需要指定可用区（AZ）的名称或者ID或者code。  可通过接口 查询可用区列表接口 获取，也可参考地区和终端节点获取。
	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}
