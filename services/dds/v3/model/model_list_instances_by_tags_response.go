package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListInstancesByTagsResponse struct {
	Instances *[]InstanceItem `json:"instances,omitempty"`

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
