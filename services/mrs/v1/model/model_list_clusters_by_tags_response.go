package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListClustersByTagsResponse struct {
	Resources *[]MrsResource `json:"resources,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListClustersByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListClustersByTagsResponse", string(data)}, " ")
}
