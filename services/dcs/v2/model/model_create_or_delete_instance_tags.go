package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateOrDeleteInstanceTags struct {
	Action string `json:"action"`

	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o CreateOrDeleteInstanceTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrDeleteInstanceTags struct{}"
	}

	return strings.Join([]string{"CreateOrDeleteInstanceTags", string(data)}, " ")
}
