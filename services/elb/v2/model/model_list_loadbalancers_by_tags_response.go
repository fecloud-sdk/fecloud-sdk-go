package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLoadbalancersByTagsResponse struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Resources      *[]ResourcesByTag `json:"resources,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListLoadbalancersByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadbalancersByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListLoadbalancersByTagsResponse", string(data)}, " ")
}
