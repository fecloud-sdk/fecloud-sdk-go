package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResourceInstance struct {
	ResourceId *string `json:"resource_id,omitempty"`

	Tags *[]TagList `json:"tags,omitempty"`

	ResourceName *string `json:"resource_name,omitempty"`
}

func (o ResourceInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstance struct{}"
	}

	return strings.Join([]string{"ResourceInstance", string(data)}, " ")
}
