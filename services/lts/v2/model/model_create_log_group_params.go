package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLogGroupParams struct {
	LogGroupName string `json:"log_group_name"`

	TtlInDays int32 `json:"ttl_in_days"`

	Tags *[]TagsBody `json:"tags,omitempty"`
}

func (o CreateLogGroupParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogGroupParams struct{}"
	}

	return strings.Join([]string{"CreateLogGroupParams", string(data)}, " ")
}
