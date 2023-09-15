package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// PageInfo 分页信息
type PageInfo struct {

	// 分页查询结果中第一条记录的ID。通常用于配合page_reverse=true查询上一页。
	PreviousMarker string `json:"previous_marker"`

	// 分页查询结果中最后一条记录的ID。通常用于查询下一页。
	NextMarker *string `json:"next_marker,omitempty"`

	// 当前的记录数。
	CurrentCount int32 `json:"current_count"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
