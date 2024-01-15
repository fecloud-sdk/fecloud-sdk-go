package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QueryImageByTagsResourceDetail struct {
	Status string `json:"status"`
}

func (o QueryImageByTagsResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryImageByTagsResourceDetail struct{}"
	}

	return strings.Join([]string{"QueryImageByTagsResourceDetail", string(data)}, " ")
}
