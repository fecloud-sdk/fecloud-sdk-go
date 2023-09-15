package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowResourcesListResponseBody struct {

	// 资源列表对象。
	Resources []ShowResourcesDetailResponseBody `json:"resources"`
}

func (o ShowResourcesListResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourcesListResponseBody struct{}"
	}

	return strings.Join([]string{"ShowResourcesListResponseBody", string(data)}, " ")
}
