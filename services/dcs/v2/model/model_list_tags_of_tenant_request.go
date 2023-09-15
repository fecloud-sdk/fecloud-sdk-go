package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListTagsOfTenantRequest Request Object
type ListTagsOfTenantRequest struct {
}

func (o ListTagsOfTenantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsOfTenantRequest struct{}"
	}

	return strings.Join([]string{"ListTagsOfTenantRequest", string(data)}, " ")
}
