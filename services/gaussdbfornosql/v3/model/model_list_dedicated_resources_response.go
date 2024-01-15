package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListDedicatedResourcesResponse struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Resources      *[]ListDedicatedResourceResult `json:"resources,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListDedicatedResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListDedicatedResourcesResponse", string(data)}, " ")
}
