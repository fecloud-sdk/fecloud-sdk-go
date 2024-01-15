package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CountEipAvailableResourcesResponse Response Object
type CountEipAvailableResourcesResponse struct {

	// - 功能说明：返回结果
	Result         *int32 `json:"result,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountEipAvailableResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountEipAvailableResourcesResponse struct{}"
	}

	return strings.Join([]string{"CountEipAvailableResourcesResponse", string(data)}, " ")
}
