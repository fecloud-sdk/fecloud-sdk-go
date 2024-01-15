package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListInstancesByResourceTagsResponse struct {
	Instances *[]InstanceResult `json:"instances,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstancesByResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesByResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesByResourceTagsResponse", string(data)}, " ")
}
