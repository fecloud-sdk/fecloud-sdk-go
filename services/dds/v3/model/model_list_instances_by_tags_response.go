package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListInstancesByTagsResponse Response Object
type ListInstancesByTagsResponse struct {

	// 实例列表。
	Instances *[]InstanceItem `json:"instances,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstancesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesByTagsResponse", string(data)}, " ")
}