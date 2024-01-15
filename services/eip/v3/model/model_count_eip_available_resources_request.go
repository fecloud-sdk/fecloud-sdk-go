package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CountEipAvailableResourcesRequest Request Object
type CountEipAvailableResourcesRequest struct {
	Body *EipResourcesAvailableV3RequestBody `json:"body,omitempty"`
}

func (o CountEipAvailableResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountEipAvailableResourcesRequest struct{}"
	}

	return strings.Join([]string{"CountEipAvailableResourcesRequest", string(data)}, " ")
}
