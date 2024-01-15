package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListTagResourceInstancesResponse struct {
	Resources *[]Resource `json:"resources,omitempty"`

	TotalCount *int32 `json:"total_count,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTagResourceInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagResourceInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListTagResourceInstancesResponse", string(data)}, " ")
}
