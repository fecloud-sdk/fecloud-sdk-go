package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Resource struct {
	ResourceDetail *interface{} `json:"resource_detail,omitempty"`

	ResourceId string `json:"resource_id"`

	ResourceName *string `json:"resource_name,omitempty"`

	Tags []Tag `json:"tags"`

	SysTags *[]Tag `json:"sys_tags,omitempty"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
