package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListInstanceTagsResult struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o ListInstanceTagsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTagsResult struct{}"
	}

	return strings.Join([]string{"ListInstanceTagsResult", string(data)}, " ")
}
